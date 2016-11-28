package bean

import (
	ssoerror "sso/error"
)

type (
	Response struct {
		Code string      `json:"code"`
		Desc string      `json:"desc"`
		Data interface{} `json:"data"`
	}
)

func NEWResponse(code string) (response *Response) {
	return &Response{Code: code, Desc: ssoerror.ErrorCodeToText(code)}
}
