package middleares

import "github.com/gin-gonic/gin"

// 中间件：请求修改
func RequestModifierMiddleware(c *gin.Context) {
	// 修改请求，例如添加额外的头信息
	c.Request.Header.Set("X-Modified-By", "RequestModifierMiddleware")

	// 执行下一个处理函数
	c.Next()
}
