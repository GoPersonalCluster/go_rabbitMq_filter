package routes

import (
	"github.com/gin-gonic/gin"

	"go-gin-redis/internal/handlers"
)

func Setup(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
) {
	api := router.Group("/api")

	{
		api.GET("/users/:id", userHandler.GetUser)
	}
}