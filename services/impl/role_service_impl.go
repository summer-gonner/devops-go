package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleServiceImpl struct {
}

func (rsi RoleServiceImpl) ListPageRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ListPageRole"})
}

func (rsi RoleServiceImpl) DetailRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DetailRole"})
}

func (rsi RoleServiceImpl) CreateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateRole"})
}
func (rsi RoleServiceImpl) DeleteRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteRole"})
}
func (rsi RoleServiceImpl) UpdateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateRole"})
}
