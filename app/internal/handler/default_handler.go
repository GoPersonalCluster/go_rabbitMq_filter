package handlers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
)
type DefaultHandler struct {
	Redis *redis.Client
}
func NewDefaultHandler(redisClient *redis.Client) *DefaultHandler {
	return &DefaultHandler{
		Redis: redisClient,
	}
}


// GetUser godoc
// @Summary      
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Success      200 {object} string
// @Router       /healthCheck [get]
func (h *DefaultHandler) HealthCheck(c *gin.Context) {
	err = h.Redis.Set(
		ctx,
		cacheKey,
		data,
		5*time.Minute,
	).Err()

	if err != nil {
		// Cache é secundário: a API continua funcionando.
	}
	c.JSON(	
		http.StatusOK,
		"API is healthy",
	)
}