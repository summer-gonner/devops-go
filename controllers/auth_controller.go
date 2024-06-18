package controllers

import (
	"devops-go/services"
	"devops-go/services/impl"
)

func (router RouterGroup) AuthController() {
	authServiceImpl := impl.AuthServiceImpl{}
	var authService services.AuthService = authServiceImpl

	authRouters := router.Group("/auth")
	{

		// @Summary 登录
		// @Description Get a list of all users
		// @Tags users
		// @Accept json
		// @Produce json
		// @Success 200 {array} User
		// @Router /auth/login [post]
		authRouters.POST("/login", authService.Login)
		authRouters.POST("/logout", authService.Logout)
	}

}
