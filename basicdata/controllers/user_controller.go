package controllers

import (
	"devops-go/services"
	"devops-go/services/impl"
)

func (router RouterGroup) UserController() {
	userServiceImpl := impl.UserServiceImpl{}
	var userService services.UserService = userServiceImpl
	userRouters := router.Group("/user")
	{
		userRouters.POST("/listPage", userService.ListPageUser)
		userRouters.POST("/create", userService.CreateUser)
		userRouters.POST("/update", userService.UpdateUser)
		userRouters.POST("/delete", userService.DeleteUser)
		userRouters.POST("/detail/:userId", userService.DetailUser)
	}
}
