package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			requestId = uuid.NewV4().String()
		}
		c.Header("X-Request-Id", requestId)
		c.Next()
	}
}
