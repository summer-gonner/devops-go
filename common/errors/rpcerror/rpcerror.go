package rpcerror

import (
	"devops-go/common/errors"
)

var _ errors.CommonError = (*RpcError)(nil)

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"error"`
}

func (e *RpcError) Error() string {
	return e.Message
}

func (e *RpcError) ErrorType() string {
	return e.Type
}

func (e *RpcError) Data() *errors.CommonErrorResp {
	return &errors.CommonErrorResp{
		Code:    e.Code,
		Message: e.Message,
		Type:    e.Type,
	}
}

// New rpc返回错误
func New(e error) error {
	msg := e.Error()[len("rpc error: code = Unknown desc = "):]
	return &RpcError{Code: errors.RpcCode, Message: msg, Type: "rpc error"}
}

// NewError 返回自定义错误，rpc返回错误
func NewError(s string, err error) error {
	msgType := err.Error()[len("rpc error: code = Unknown desc = "):]
	return &RpcError{Code: errors.RpcCode, Message: s, Type: msgType}
}
