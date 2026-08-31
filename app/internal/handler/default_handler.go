package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Success      200 {object} string
// @Router       /healthCheck [get]
func HealthCheck(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		"API is healthy",
	)
}
