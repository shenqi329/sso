package error

import (
	"errors"
	"fmt"
)

var (
	//通用错误
	ErrorIllegalParams       = NEWError(CommonIllegalParams)
	ErrorResourceExist       = NEWError(CommonResourceExist)
	ErrorNotFound            = NEWError(CommonResourceNoExist)
	ErrorInternalServerError = NEWError(CommonInternalServerError)

	//注册
	ErrorRegisterEmailInUse = NEWError(RegisterEmailInUse)
)

const (
	//通用状态模块码 [000]
	CommonSuccess             = "00000001"
	CommonIllegalParam        = "00000002"
	CommonResourceNoExist     = "00000003"
	CommonResourceExist       = "00000004"
	CommonInternalServerError = "00000005"

	//注册
	RegisterEmailInUse = "00001001"
)

var codeText = map[string]string{
	//通用状态
	CommonSuccess:             "success",
	CommonIllegalParam:        "illegal parameter",
	CommonResourceNoExist:     "resource doesn't exist",
	CommonResourceExist:       "resource already exists",
	CommonInternalServerError: "internal server error",

	//注册
	RegisterEmailInUse: "the email address is currently in use",
}

func ErrorCodeToText(code string) string {
	return codeText[code]
}

type (
	SSOError struct {
		Code string
		Desc string
	}
)

func NEWError(code string) *SSOError {
	return &SSOError{Code: code, Desc: ErrorCodeToText(code)}
}

func (err *SSOError) Error() {
	fmt.Sprintf("code = %s,desc = %s", err.Code, err.Desc)
}
