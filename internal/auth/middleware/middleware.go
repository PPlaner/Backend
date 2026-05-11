package middleware

import (
	"net/http"
	"strings"

	"github.com/PPlaner/Backend/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "invalid authorization header")
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "invalid token claims")
			c.Abort()
			return
		}

		userIDValue, exists := claims["user_id"]
		if !exists {
			response.Error(c, http.StatusUnauthorized, "user id not found in token")
			c.Abort()
			return
		}

		userIDFloat, ok := userIDValue.(float64)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "invalid user id type")
			c.Abort()
			return
		}

		userID := int(userIDFloat)
		c.Set("user_id", userID)

		c.Next()
	}
}
