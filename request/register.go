package request

import ()

type (
	Register struct {
		UserName   string `json:"username,omitempty"`
		Email      string `json:"email,omitempty"`
		VerifyCode string `json:"code,omitempty"`
		Password   string `json:"password,omitempty"`
	}
)
