package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/auth"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/auth/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMiddlewareTest() {
	gin.SetMode(gin.TestMode)
}
func TestJWTMiddleware_ValidToken(t *testing.T) {
	setupMiddlewareTest()

	token, err := auth.GenerateToken(
		123,
		"user@example.com",
		"admin",
	)

	require.NoError(t, err)

	router := gin.New()

	router.GET(
		"/protected",
		middleware.JWTMiddleware(),
		func(c *gin.Context) {

			c.Status(http.StatusOK)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	rec := httptest.NewRecorder()

	// Executa efetivamente o middleware + handler
	router.ServeHTTP(rec, req)
	println(rec.Code == http.StatusOK)
	// Agora sim verificamos a resposta produzida
	assert.Equal(t, 200, rec.Code)
}
