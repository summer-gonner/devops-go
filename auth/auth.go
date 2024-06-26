package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest 定义登录请求的结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 添加登录接口的路由处理
	router.POST("/login", handleLogin)

	// 启动服务，并监听端口
	port := ":8081"
	router.Run(port)
}

// handleLogin 是处理登录请求的处理函数
func handleLogin(c *gin.Context) {
	var loginReq LoginRequest

	// 解析请求中的 JSON 数据到 loginReq 结构体
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 假设这里是验证用户名密码的逻辑
	// 这里简单模拟，实际应用中需要进行安全验证和业务逻辑处理
	if loginReq.Username == "admin" && loginReq.Password == "password" {
		// 登录成功返回响应
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		// 登录失败返回错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
