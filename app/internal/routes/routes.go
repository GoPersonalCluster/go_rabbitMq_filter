package routes

import (
	handlers "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/handler"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/healthCheck", handlers.HealthCheck)
	}

}
