package routes

import (
	middleware "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/auth/middleware"
	handlers "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/handler"
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
