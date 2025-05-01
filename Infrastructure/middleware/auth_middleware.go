package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("your-secret-key-here")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Log the token for debugging
		fmt.Printf("Token being validated: %s\n", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Log the signing method
			fmt.Printf("Token signing method: %v\n", token.Method)
			return secretKey, nil
		})

		if err != nil {
			fmt.Printf("Token validation error: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("Token is not valid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		// Log the claims
		claims := token.Claims.(jwt.MapClaims)
		fmt.Printf("Token claims: %+v\n", claims)

		c.Set("username", claims["username"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
