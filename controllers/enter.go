package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	routerGroup := r.Group("/api")
	routerGroupApp := RouterGroup{routerGroup}

	routerGroupApp.AuthController()
	routerGroupApp.UserController()
	routerGroupApp.RoleController()
	return r
}
