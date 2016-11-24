package request

import ()

type (
	Register struct {
		UserName string `json:"username,omitempty" form:"username"`
		Email    string `json:"email,omitempty" form:"email"`
		Code     string `json:"code,omitempty" form:"code"`
		Password string `json:"password,omitempty" form:"password"`
	}
)
