package bean

import ()

type (
	Login struct {
		UserName string `json:"username,omitempty" form:"username"`
		Password string `json:"password,omitempty" form:"password"`
		DeviceId string `json:"deviceId,omitempty" form:"deviceId"`
		AppId    string `json:"appId,omitempty" form:"appId"`
		Platform string `json:"platform,omitempty" form:"platform"`
	}
)
