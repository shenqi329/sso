package bean

type (
	Response struct {
		Code string      `json:"code"  form:"code"`
		Desc string      `json:"desc"  form:"desc"`
		Data interface{} `json:"data"  form:"data"`
	}
)
