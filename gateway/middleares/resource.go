package middleares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 示例处理函数，模拟获取资源
func getResourceHandler(c *gin.Context) {
	// 从数据库或其他服务获取资源
	data := "Some resource data"
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
