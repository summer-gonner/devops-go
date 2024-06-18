package services

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	ListPageUser(c *gin.Context)
	DetailUser(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}
