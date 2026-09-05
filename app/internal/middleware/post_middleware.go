package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func PostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		// Executa o handler e os middlewares seguintes
		c.Next()

		// A partir daqui o processamento voltou do handler
		duration := time.Since(start)

		// Exemplo de header adicionado após o processamento
		c.Header("X-Processing-Time", duration.String())
	}
}
