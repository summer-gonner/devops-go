package errors

type CommonError interface {
	Error() string
	ErrorType() string

	Data() *CommonErrorResp
}

type CommonErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"error"`
}
