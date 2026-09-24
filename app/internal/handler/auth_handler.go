package handlers

import (
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/handler_model"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/gin-gonic/gin"
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
	var body handler_model.Authentication

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := db.GetDbConnection()
	username, err := vo.NewUsername(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}
	password, err := vo.NewPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	var existingUser postgresql_entity.User

	validation := db.Where(&postgresql_entity.User{
		Username: username,
		Password: password,
	}).First(&existingUser).Error

}
