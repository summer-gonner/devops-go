package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServiceImpl struct {
}

func (usi UserServiceImpl) ListPageUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ListPageUser"})
}

func (usi UserServiceImpl) DetailUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DetailUser"})
}

func (usi UserServiceImpl) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateUser"})
}
func (usi UserServiceImpl) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteUser"})
}
func (usi UserServiceImpl) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateUser"})
}
