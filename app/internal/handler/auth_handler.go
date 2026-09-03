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
// @Success      401 {object} string
// @Router       /healthCheck [post]
func Authentication(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		"auth request",
	)
}
