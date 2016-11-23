package error

import (
	//"errors"
	"fmt"
)

var (
	//通用错误
	ErrorIllegalParams       = NEWError(CommonIllegalParams)
	ErrorResourceExist       = NEWError(CommonResourceExist)
	ErrorNotFound            = NEWError(CommonResourceNoExist)
	ErrorInternalServerError = NEWError(CommonInternalServerError)
	ErrorTokenInvalidated    = NEWError(CommonTokenInvalidated)

	//注册
	ErrorRegisterEmailInUse      = NEWError(RegisterEmailInUse)
	ErrorRegisterEmailFormat     = NEWError(RegisterEmailFormatError)
	ErrorRegisterErrorCode       = NEWError(RegisterErrorCode)
	ErrorRegisterUserExist       = NEWError(RegisterUserExist)
	ErrorUserNameFormatError     = NEWError(RegisterUserNameFormatError)
	ErrorPasswordFormatError     = NEWError(RegisterPasswordFormatError)
	ErrorUserNameOrPasswordFail  = NEWError(RegisterUserNameOrPasswordFail)
	ErrorSameOriginalNewPassword = NEWError(RegisterSameOriginalNewPassword)
	ErrorPasswordWrong           = NEWError(RegisterPasswordWrong)
)

const (
	//通用状态模块码 [000]
	CommonSuccess             = "00000001"
	CommonIllegalParams       = "00000002"
	CommonResourceNoExist     = "00000003"
	CommonResourceExist       = "00000004"
	CommonInternalServerError = "00000005"
	CommonTokenInvalidated    = "00000006"

	//注册
	RegisterEmailInUse              = "00001001"
	RegisterEmailFormatError        = "00001002"
	RegisterErrorCode               = "00001003"
	RegisterUserExist               = "00001004"
	RegisterUserNameFormatError     = "00001005"
	RegisterPasswordFormatError     = "00001006"
	RegisterUserNameOrPasswordFail  = "00001007"
	RegisterSameOriginalNewPassword = "00001008"
	RegisterPasswordWrong           = "00001009"
)

var codeText = map[string]string{
	//通用状态
	CommonSuccess:             "success",
	CommonIllegalParams:       "illegal parameter",
	CommonResourceNoExist:     "resource doesn't exist",
	CommonResourceExist:       "resource already exists",
	CommonInternalServerError: "internal server wrong",
	CommonTokenInvalidated:    "token invalidated",

	//注册
	RegisterEmailInUse:              "the email address is currently in use",
	RegisterEmailFormatError:        "the email format wrong",
	RegisterErrorCode:               "wrong verify code",
	RegisterUserExist:               "user already exists",
	RegisterUserNameFormatError:     "username format wrong",
	RegisterPasswordFormatError:     "password format wrong",
	RegisterUserNameOrPasswordFail:  "username or passowd wrong",
	RegisterSameOriginalNewPassword: "original password is same with new password",
	RegisterPasswordWrong:           "password wrong",
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

func (err *SSOError) Error() string {
	errString := fmt.Sprintf("code = %s,desc = %s", err.Code, err.Desc)
	return errString
}
