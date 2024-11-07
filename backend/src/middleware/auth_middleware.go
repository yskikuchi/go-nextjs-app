package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/service"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	headerParts := strings.Split(authHeader, " ")

	if len(headerParts) != 2 {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	if headerParts[0] != "Bearer" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	tokenString := headerParts[1]

	token, err := service.ParseToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
