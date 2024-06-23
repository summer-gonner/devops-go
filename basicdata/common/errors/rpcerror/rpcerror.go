package rpcerror

import (
	errors2 "devops-go/basicdata/common/errors"
)

var _ errors2.CommonError = (*RpcError)(nil)

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

func (e *RpcError) Data() *errors2.CommonErrorResp {
	return &errors2.CommonErrorResp{
		Code:    e.Code,
		Message: e.Message,
		Type:    e.Type,
	}
}

// New rpc返回错误
func New(e error) error {
	msg := e.Error()[len("rpc error: code = Unknown desc = "):]
	return &RpcError{Code: errors2.RpcCode, Message: msg, Type: "rpc error"}
}

// NewError 返回自定义错误，rpc返回错误
func NewError(s string, err error) error {
	msgType := err.Error()[len("rpc error: code = Unknown desc = "):]
	return &RpcError{Code: errors2.RpcCode, Message: s, Type: msgType}
}
