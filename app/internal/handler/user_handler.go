package handlers

import (
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/handler_model"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUser godoc
// @Summary      Creates a user account
// @Description  Creates a user account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body User true "User data"
// @Success      201 {object} User
// @Failure      400 {object} string
// @Failure      409 {object} string
// @Failure      500 {object} string
// @Router       /user [post]
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

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request body",
			},
		)
		return
	}

	var existingUser postgresql_entity.User

	result := db.Where("email = ?", user.Email.Value()).First(&existingUser).Error

	if result.Error() == "" {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": "email already registered",
			},
		)
		return
	}

	if result.Error() != gorm.ErrRecordNotFound.Error() {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "failed to check email",
			},
		)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "failed to create user",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		user,
	)
}
