package blademaster

import (
	"github.com/zy-free/micro-study/lib/ecode"
	"net/http"
)

type json struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Render(code int, data interface{}, err error) (int, interface{}) {
	httpCode := http.StatusOK
	if code != 0 {
		httpCode = code
	}
	bcode := ecode.Cause(err)
	// TODO app allow 5xx?
	/*
		if bcode.Code() == -500 {
			code = http.StatusServiceUnavailable
		}
	*/
	return httpCode, json{
		Code:    bcode.Code(),
		Message: bcode.Message(),
		Data:    data,
	}
}
