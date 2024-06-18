package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func isValidToken(token string) bool {
	// 在实际情况下，这里应该是你的具体的 Token 验证逻辑
	// 例如，验证 Token 是否在数据库或缓存中有效
	validToken := "my_secret_token"
	return token == validToken
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if !isValidToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token 验证通过，继续处理请求
		c.Next()
	}
}
