package routes

import (
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	handlers "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/handler"
	middleware "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, cache *cache.RedisCache) {
	SetupUnsafeRoutes(router)
	SetupProtectedRoutes(router, cache)
}

func SetupUnsafeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/healthCheck", handlers.HealthCheck)
		api.GET("/Authentication", handlers.Authentication)
		api.POST("/User", handlers.CreateUser)

	}
}

func SetupProtectedRoutes(router *gin.Engine, cache *cache.RedisCache) {

	protected := router.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware(), middleware.RedisHealthMiddleware(cache))
	{
		protected.GET("/GetPipelineLog", handlers.GetPipelineLog)
	}

}
