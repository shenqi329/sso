package util

//状态码解释  [00]           [000]        [000]
//          前两位表示系统 中间三位表示模块 后三位表示不同状态
//          系统码          模块码         状态码

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

var statusText = map[string]string{
	//通用状态
	CommonSuccess:             "success",
	CommonIllegalParam:        "illegal parameter",
	CommonResourceNoExist:     "resource doesn't exist",
	CommonResourceExist:       "resource already exists",
	CommonInternalServerError: "internal server error",

	//注册
	RegisterEmailInUse: "the email address is currently in use",
}

func StatusText(code string) string {
	return statusText[code]
}
