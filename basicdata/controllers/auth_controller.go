package controllers

import (
	"devops-go/basicdata/services"
	"devops-go/basicdata/services/impl"
)

// AuthController handles authentication routes
func (router RouterGroup) AuthController() {
	authServiceImpl := impl.AuthServiceImpl{}
	var authService services.AuthService = authServiceImpl

	authRouters := router.Group("/auth")
	{

		authRouters.POST("/login", authService.Login)

		authRouters.POST("/logout", authService.Logout)
	}
}
