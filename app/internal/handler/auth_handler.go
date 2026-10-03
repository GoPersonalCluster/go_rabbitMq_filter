package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/handler_model"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	redismodels "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/gin-gonic/gin"
)

type authDTO struct {
	username vo.Username
	password vo.Password
}

// GetUser godoc
// @Summary
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Success      200 {object} string
// @Failure      401 {object} string
// @Router       /api/v1/Authentication [post]
func Authentication(c *gin.Context) {

	// cache := cache.NewRedisCache()
	// authIp, err := cache.Get(c, c.ClientIP())

	db := db.GetDbConnection()

}

func ensureBodyContentIsValid(c *gin.Context, dto *authDTO) (*authDTO, error) {
	var body handler_model.Authentication

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return dto, fmt.Errorf("no body content was detected")
	}

	username, err := vo.NewUsername(body.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return dto, fmt.Errorf("invalid username format")
	}
	password, err := vo.NewPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return dto, fmt.Errorf("invalid password format")
	}
	dto.username = username
	dto.password = password

	return dto, nil
}
func authenticateUser() {
	var existingUser postgresql_entity.User

	validation := db.Where(&postgresql_entity.User{
		Username: username,
		Password: password,
	}).First(&existingUser).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}
}
func dataComparingStep(c *gin.Context) {
	cache := cache.NewRedisCache()
	var err = authenticateIp(cache, c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

}

func authenticateIp(c *cache.RedisCache, ctx *gin.Context) error {
	authIp, err := c.GetBytes(
		ctx,
		ctx.ClientIP(),
	)
	if err != nil {
		return err
	}

	data := &redismodels.AuthenticationIP{}
	json.Unmarshal(authIp, &data)

	if data.Score <= 0 {
		return fmt.Errorf("user score is below zero")
	}

	entity := redismodels.NewAuthenticationIp(net.ParseIP(ctx.ClientIP()))
	jsonb, err := json.Marshal(entity)

	cache.NewRedisCache().SetBytes(ctx, ctx.ClientIP(), jsonb, 360)
	return nil
}
