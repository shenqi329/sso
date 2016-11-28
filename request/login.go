package request

import ()

type (
	Login struct {
		UserName string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		DeviceId string `json:"deviceId,omitempty"`
		AppId    string `json:"appId,omitempty"`
		Platform string `json:"platform,omitempty"`
	}
)
