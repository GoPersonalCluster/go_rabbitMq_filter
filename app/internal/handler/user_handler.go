package handlers

import (
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/handler_model"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	"github.com/gin-gonic/gin"
)

// swagger annotation
// CreateUser godoc
// @Summary Creates a user account
// @Description Creates a user account
// @Tags User
// @Accept json
// @Produce json
// @Param user body handler_model.CreateUser true "User data"
// @Success 201 {object} postgresql_entity.User
// @Failure 400 {string} string
// @Failure 409 {string} string
// @Failure 500 {string} string
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
	db := db.GetDbConnection()

	var body handler_model.CreateUser
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := postgresql_entity.NewUser(
		body.Username,
		body.Email,
		body.Password,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request body",
			},
		)
		return
	}

	var existingUser postgresql_entity.User

	validation := db.Where("email = ?", user.Username.Value).First(&existingUser).Error

	if validation == nil && existingUser.ID != 0 {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": "Invalid email or username",
			},
		)
		return
	}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request body",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		user,
	)
}
