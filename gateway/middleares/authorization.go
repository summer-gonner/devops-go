package middleares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件：授权检查
func AuthorizationMiddleware(c *gin.Context) {
	// 模拟授权逻辑
	token := c.GetHeader("Authorization")
	if token != "Bearer token123" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 执行下一个处理函数
	c.Next()
}
