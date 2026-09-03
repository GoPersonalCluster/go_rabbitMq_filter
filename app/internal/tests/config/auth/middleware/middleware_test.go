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

func TestJWTMiddleware_MissingAuthorizationHeader(t *testing.T) {
	setupMiddlewareTest()

	router := gin.New()
	router.Use(middleware.JWTMiddleware())

	called := false

	router.GET("/protected", func(c *gin.Context) {
		called = true
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
	assert.Contains(t, rec.Body.String(), "authorization header is required")
}

func TestJWTMiddleware_InvalidAuthorizationScheme(t *testing.T) {
	setupMiddlewareTest()

	router := gin.New()
	router.Use(middleware.JWTMiddleware())

	called := false

	router.GET("/protected", func(c *gin.Context) {
		called = true
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Basic abc123",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
	assert.Contains(
		t,
		rec.Body.String(),
		"authorization header must use Bearer scheme",
	)
}

func TestJWTMiddleware_MalformedAuthorizationHeader(t *testing.T) {
	setupMiddlewareTest()

	router := gin.New()
	router.Use(middleware.JWTMiddleware())

	called := false

	router.GET("/protected", func(c *gin.Context) {
		called = true
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
}

func TestJWTMiddleware_InvalidToken(t *testing.T) {
	setupMiddlewareTest()

	router := gin.New()
	router.Use(middleware.JWTMiddleware())

	called := false

	router.GET("/protected", func(c *gin.Context) {
		called = true
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer invalid-token",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
	assert.Contains(
		t,
		rec.Body.String(),
		"invalid or expired token",
	)
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
	router.Use(middleware.JWTMiddleware())

	router.GET("/protected", func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		require.True(t, exists)

		email, exists := c.Get("email")
		require.True(t, exists)

		role, exists := c.Get("role")
		require.True(t, exists)

		assert.Equal(t, int64(123), userID)
		assert.Equal(t, "user@example.com", email)
		assert.Equal(t, "admin", role)

		c.Status(http.StatusOK)
	})

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

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestJWTMiddleware_TokenWithBearerCaseInsensitive(t *testing.T) {
	setupMiddlewareTest()

	token, err := auth.GenerateToken(
		123,
		"user@example.com",
		"admin",
	)

	require.NoError(t, err)

	router := gin.New()
	router.Use(middleware.JWTMiddleware())

	router.GET("/protected", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	testCases := []string{
		"Bearer " + token,
		"bearer " + token,
		"BEARER " + token,
		"BeArEr " + token,
	}

	for _, authorization := range testCases {
		t.Run(authorization[:6], func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodGet,
				"/protected",
				nil,
			)

			req.Header.Set(
				"Authorization",
				authorization,
			)

			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
