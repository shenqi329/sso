package request

import ()

type (
	Update struct {
		Name     string `json:"name,omitempty" form:"name"` //姓名
		Icon     string `json:"icon,omitempty" form:"icon"` //头像图片url
		Sex      string `json:"sex,omitempty" form:"sex"`   //性别
		NickName string `json:"nickname,omitempty" form:"nickname"`
		Birthday *int64 `json:"birthday,omitempty" form:"birthday"` //生日
	}
)
