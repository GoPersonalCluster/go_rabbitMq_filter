package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary
// @Description  Creates a user account
// @Tags         User
// @Produce      json
// @Success      201 {object} string
// @Error      400 {object} string
// @Error      501 {object} string
// @Router       /user [post]
func CreateUser(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		"API Log Endpoint",
	)
}
