package handlers

import (
	"encoding/json"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// GetUser godoc
// @Summary
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Success      200 {object} string
// @Failure      401 {object} string
// @Router       /api/v1/Authentication [post]
func Authentication(c *gin.Context) {
	// var body handler_model.Authentication

	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// // cache := cache.NewRedisCache()
	// // authIp, err := cache.Get(c, c.ClientIP())

	// db := db.GetDbConnection()
	// username, err := vo.NewUsername(body.Username)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid username or password",
	// 	})
	// 	return
	// }
	// password, err := vo.NewPassword(body.Password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid username or password",
	// 	})
	// 	return
	// }

	// cache := cache.NewRedisCache()

	// var existingUser postgresql_entity.User

	// validation := db.Where(&postgresql_entity.User{
	// 	Username: username,
	// 	Password: password,
	// }).First(&existingUser).Error

}
func AuthenticateIp(c *cache.RedisCache, ctx *gin.Context) (
	*redismodels.AuthenticationIP, error) {
	authIp, err := c.GetBytes(
		ctx,
		ctx.ClientIP(),
	)
	if err != nil {
		return nil, err
	}

	data := &redismodels.AuthenticationIP{}
	json.Unmarshal(authIp, &data)

	if data.IP != "" {
		return data, nil
	}
	entity := redismodels.NewAuthenticationIp(ctx.ClientIP())

}
