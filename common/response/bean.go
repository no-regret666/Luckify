package response

import "net/http"

type SuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type NullJson struct{}

func Success(data interface{}) *SuccessBean {
	return &SuccessBean{
		http.StatusOK,
		"OK",
		data,
	}
}

type ErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(code uint32, msg string) *ErrorBean {
	return &ErrorBean{
		code,
		msg,
	}
}
