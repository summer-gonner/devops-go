package services

import (
	"github.com/gin-gonic/gin"
)

type RoleService interface {
	ListPageRole(c *gin.Context)
	DetailRole(c *gin.Context)
	CreateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
	UpdateRole(c *gin.Context)
}
