package controllers

import (
	"devops-go/services"
	"devops-go/services/impl"
)

func (router RouterGroup) RoleController() {
	roleServiceImpl := impl.RoleServiceImpl{}
	var roleService services.RoleService = roleServiceImpl
	roleRouters := router.Group("/role")
	{
		roleRouters.POST("/listPage", roleService.ListPageRole)
		roleRouters.POST("/create", roleService.CreateRole)
		roleRouters.POST("/update", roleService.UpdateRole)
		roleRouters.POST("/delete", roleService.DeleteRole)
		roleRouters.POST("/detail/:roleId", roleService.DetailRole)
	}
}
