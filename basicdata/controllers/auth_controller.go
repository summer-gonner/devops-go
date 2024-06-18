package controllers

import (
	"devops-go/services"
	"devops-go/services/impl"
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
