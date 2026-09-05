package routes

import (
	handlers "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/handler"
	middleware "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	SetupUnsafeRoutes(router)
	SetupProtectedRoutes(router)
}

func SetupUnsafeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/healthCheck", handlers.HealthCheck)
	}
}

func SetupProtectedRoutes(router *gin.Engine) {

	protected := router.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware())
	{
		//protected.GET("/profile", handlers.GetProfile)
	}

}
