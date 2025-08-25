package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles Cross-Origin Resource Sharing (CORS) headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow origins - you can specify specific domains instead of "*" for production
		c.Header("Access-Control-Allow-Origin", "*")
		
		// Allow methods
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		
		// Allow headers
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		
		// Allow credentials
		c.Header("Access-Control-Allow-Credentials", "true")
		
		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}