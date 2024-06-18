package services

import "github.com/gin-gonic/gin"

type AuthService interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	CheckToken(c *gin.Context)
}
