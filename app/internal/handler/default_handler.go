package handlers

import (
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/di"
	"github.com/gin-gonic/gin"
)

var dbConfig = di.NewDbConnectionDi()

func Init() *di.DbConnectionDi {
	if dbConfig == nil {
		dbConfig = di.NewDbConnectionDi()
	}

	return dbConfig
}

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
