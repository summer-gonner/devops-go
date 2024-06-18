package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"` //0 成功  1失败
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(data any, msg string, c *gin.Context) {

	Result(SUCCESS, data, msg, c)
}

func SuccessWithoutMsg(data any, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(FAIL, data, msg, c)
}
func FailWithoutMsg(data any, c *gin.Context) {
	Result(FAIL, data, "操作失败", c)
}
