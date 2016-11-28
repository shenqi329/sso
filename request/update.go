package request

import ()

type (
	Update struct {
		Name     string `json:"name,omitempty"`     //姓名
		Icon     string `json:"icon,omitempty"`     //头像图片url
		Sex      string `json:"sex,omitempty"`      //性别
		NickName string `json:"nickname,omitempty"` //别名
		Birthday *int64 `json:"birthday,omitempty"` //生日
	}
)
