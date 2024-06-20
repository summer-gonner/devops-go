package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestServiceImpl struct {
}

func (rsi TestServiceImpl) Trace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "回传成功", "code": 1})
}
