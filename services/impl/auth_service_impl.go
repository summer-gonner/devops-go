package impl

import (
	"devops-go/common/res"
	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
}

func (asi AuthServiceImpl) Login(c *gin.Context) {

	res.SuccessWithoutMsg("登录成功", c)
}

func (asi AuthServiceImpl) Logout(c *gin.Context) {
	res.SuccessWithoutMsg("退出登录成功", c)

}
func (asi AuthServiceImpl) CheckToken(c *gin.Context) {
	res.SuccessWithoutMsg("token生效", c)
}
