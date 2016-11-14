package util

const (
	StatusOK              = "000001"
	StatusIllegalParam    = "000002"
	StatusResourceNoExist = "000003"
	StatusResourceExist   = "000004"
)

var statusText = map[string]string{
	StatusOK:              "success",
	StatusIllegalParam:    "illegal parameter",
	StatusResourceNoExist: "resource no exist",
	StatusResourceExist:   "resource exist",
}

func StatusText(code string) string {
	return statusText[code]
}

type SSOError struct {
	Code    string
	Message string
}

func (e *SSOError) Error() string {
	return e.Message
}
