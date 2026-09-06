package middleware

import (
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"

	"github.com/gin-gonic/gin"
)

func RedisHealthMiddleware(
	redisCache *cache.RedisCache,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		if err := redisCache.Ping(
			c.Request.Context(),
		); err != nil {

			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "redis unavailable",
			})

			c.Abort()

			return
		}

		c.Next()
	}
}
