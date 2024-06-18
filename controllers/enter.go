package controllers

import (
	"devops-go/common/res"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// 设置全局异常处理中间件
	r.Use(func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
				// 构造异常返回体
				// 返回异常信息给客户端
				res.FailWithoutMsg(r, c)
				c.Abort()
			}
		}()
		c.Next()
	})
	routerGroup := r.Group("/api")
	routerGroupApp := RouterGroup{routerGroup}

	routerGroupApp.AuthController()
	routerGroupApp.UserController()
	routerGroupApp.RoleController()
	return r
}
