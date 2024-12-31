package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/todos-api/jovi345/formatter"
	"github.com/todos-api/jovi345/token"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			msg := formatter.SendResponse("Failed", "Authorization header is missing")
			c.JSON(http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			msg := formatter.SendResponse("Failed", "Invalid authorization header")
			c.JSON(http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		accessToken := tokenParts[1]
		secretKey := os.Getenv("JWT_ACCESS_KEY")

		claims, err := token.ValidateToken(accessToken, secretKey)
		if err != nil {
			msg := formatter.SendResponse("Failed", "Invalid or expired token")
			c.JSON(http.StatusForbidden, msg)
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
