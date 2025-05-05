package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var jwtService *JWTService

func SetJWTService(service *JWTService) {
	jwtService = service
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if jwtService == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT service not initialized"})
			c.Abort()
			return
		}

		// Validate token using JWT service
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			fmt.Printf("Token validation error: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

//
