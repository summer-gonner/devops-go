package errorx

import (
	errors2 "devops-go/basicdata/common/errors"
)

var _ errors2.CommonError = (*ErrorX)(nil)

type ErrorX struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"error"`
}

func (e *ErrorX) Error() string {
	return e.Message
}

func (e *ErrorX) ErrorType() string {
	return e.Type
}

func (e *ErrorX) Data() *errors2.CommonErrorResp {
	return &errors2.CommonErrorResp{
		Code:    e.Code,
		Message: e.Message,
		Type:    e.Type,
	}
}

func New(s string) error {
	return &ErrorX{Code: errors2.BaseCode, Message: s, Type: "base error"}
}

func NewCodeErr(code int, s string) error {
	return &ErrorX{Code: code, Message: s, Type: "base error"}
}
