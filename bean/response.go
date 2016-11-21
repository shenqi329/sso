package bean

import (
	ssoerror "sso/error"
)

type (
	Response struct {
		Code string      `json:"code"  form:"code"`
		Desc string      `json:"desc"  form:"desc"`
		Data interface{} `json:"data"  form:"data"`
	}
)

func NEWResponse(code string) (response *Response) {
	return &Response{Code: code, Desc: ssoerror.ErrorCodeToText(code)}
}
