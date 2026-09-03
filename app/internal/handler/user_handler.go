package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary
// @Description  Creates a user account
// @Tags         test
// @Produce      json
// @Success      201 {object} string
// @Success      400 {object} string
// @Router       /user [post]
func User(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		"API Log Endpoint",
	)
}
