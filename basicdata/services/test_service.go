package services

import "github.com/gin-gonic/gin"

type TestService interface {
	Trace(c *gin.Context)
}
