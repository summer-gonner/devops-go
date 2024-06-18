package impl

import (
	"devops-go/common/res"
	"devops-go/mappers"
	"devops-go/models/request"
	"devops-go/models/vo"
	"github.com/gin-gonic/gin"
)

var authServiceInstance *AuthServiceImpl

func InitAuthServerImpl() *AuthServiceImpl {
	if authServiceInstance == nil {
		authMapper := mappers.AuthMapperSql{} // 使用具体的实现类来初始化 authMapper
		authServiceInstance = &AuthServiceImpl{
			authMapper: &authMapper, // 使用指针类型来确保不是 nil
		}
	}
	return authServiceInstance
}

type AuthServiceImpl struct {
	authMapper mappers.AuthMapper
}

// @Summary 登录
// @Description 登录接口
// @Tags auth
// @Accept json
// @Produce json
// @Param requestBody body request.LoginRequest true "登录请求体"
// @Success 200 {object} vo.LoginVo
// @Router /auth/login [post]
func (asi AuthServiceImpl) Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		res.Fail("", "Invalid JSON format", c)
		return
	}
	loginUser := InitAuthServerImpl().authMapper.QueryUserByUsernameAndPassword(loginRequest.Username, loginRequest.Password)
	if loginUser == nil && loginUser.Id == 0 {
		res.Fail("", "用户名或者密码不对", c)
		return
	} else {
		var loginVo vo.LoginVo
		loginVo.Username = loginUser.Username
		loginVo.Password = loginUser.Password
		res.SuccessWithoutMsg(loginVo, c)
	}

}

// @Summary 退出登录
// @Description 退出登录接口
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/logout [post]
func (asi AuthServiceImpl) Logout(c *gin.Context) {
	res.SuccessWithoutMsg("退出登录成功", c)

}

// @Summary 检查 Token
// @Description 检查 Token 是否有效
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/checkToken [get]
func (asi AuthServiceImpl) CheckToken(c *gin.Context) {
	res.SuccessWithoutMsg("token生效", c)
}
