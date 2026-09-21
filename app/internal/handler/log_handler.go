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
// @Router       /api/v1/getPipelineLog [get]
func GetPipelineLog(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		"API Log Endpoint",
	)
}
