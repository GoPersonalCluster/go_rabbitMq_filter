package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/handler_model"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	redismodels "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authDTO struct {
	username                vo.Username
	password                vo.Password
	bodyContentError        error
	user                    postgresql_entity.User
	authUserError           error
	redisIp                 redismodels.AuthenticationIP
	authIpError             error
	IP                      string
	authenticationUserGroup postgresql_entity.AutherticationAccountGroupsUser
	authenticationGroup     postgresql_entity.AuthenticationAccountGroups
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
	db := db.GetDbConnection()
	cache := cache.NewRedisCache()

	dto := &authDTO{}

	dto = ensureBodyContentIsValid(c, dto)
	dto = authenticateUser(db, dto)
	dto = authenticateIp(cache, c, dto)

}

func ensureBodyContentIsValid(c *gin.Context, dto *authDTO) *authDTO {
	var body handler_model.Authentication

	if err := c.ShouldBindJSON(&body); err != nil {
		dto.bodyContentError = fmt.Errorf("no body content was detected")
		return dto
	}

	username, err := vo.NewUsername(body.Username)
	if err != nil {
		dto.bodyContentError = fmt.Errorf("invalid username format")
		return dto
	}
	password, err := vo.NewPassword(body.Password)
	if err != nil {
		dto.bodyContentError = fmt.Errorf("invalid password format")
		return dto
	}
	dto.username = username
	dto.password = password

	return dto
}
func authenticateUser(db *gorm.DB, dto *authDTO) *authDTO {
	var existingUser *postgresql_entity.User

	err := db.Where(&postgresql_entity.User{
		Username: dto.username,
		Password: dto.password,
	}).First(&existingUser).Error

	dto.authUserError = err
	dto.user = dto.user

	return dto
}

func dataComparingStep(c *gin.Context, dto *authDTO) *authDTO {
	db := db.GetDbConnection()
	cache := cache.NewRedisCache()

	if dto.authUserError && dto.authIpError {

	}

	// switch {
	// 	case dto.authUserError != nil && dto.redisIp.Score > 0:
	// 	case dto.authUserError != nil && dto.redisIp.Score > 0:
	// 	case
	// }

}

func dataComparingStepSpam() {

}

func dataComparingStepValidarGruposUsuario(c *cache.RedisCache,
	ctx *gin.Context,
	dto *authDTO,
	db *gorm.DB) *authDTO {

	entity := redismodels.NewAuthenticationIp()
	dto.IP = ctx.ClientIP()

	jsonb, err := json.Marshal(entity)
	if err != nil {
		dto.authIpError = err
	}

	cache.NewRedisCache().
		SetBytes(ctx, ctx.ClientIP(), jsonb, 360)

	dto.redisIp = *entity

	var bl postgresql_entity.AuthenticationIpBlacklist
	validation := db.Where("IP = ?", dto.IP).First(&bl).Error
	if validation != nil {
		dto.authIpError = fmt.Errorf("ip is blacklisted")
		return dto
	}
	var validationGroup postgresql_entity.AuthenticationAccountGroups

	vg := db.Where("IP = ?", dto.IP).First(&validationGroup)

	if vg.Error != nil && vg != nil {

		dto.authenticationGroup = validationGroup

		var validationGroupUser postgresql_entity.AutherticationAccountGroupsUser
		vgu := db.Where("userId = ?", dto.user.ID).First(&validationGroupUser)

		if vgu.Error != nil && vgu != nil {
			validationGroupUser = validationGroupUser
			if validationGroupUser.Enabled && 

		}



	}

	return dto
}
