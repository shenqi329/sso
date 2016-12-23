package bean

import (
	"time"
)

const (
	UserSexMale   = "m"
	UserSexFemale = "f"
)

type (
	User struct {
		ID               int64      `xorm:"'id'" json:"id,omitempty"`
		UserName         string     `xorm:"'username'" json:"username,omitempty"`
		Password         string     `xorm:"'password'" json:"password,omitempty"`
		Name             string     `xorm:"'name'" json:"name,omitempty"` //姓名
		Icon             string     `xorm:"'icon'" json:"icon,omitempty"` //头像图片url
		Sex              string     `xorm:"'sex'" json:"sex,omitempty"`   //性别
		NickName         string     `xorm:"'nickname'" json:"nickname,omitempty"`
		Email            string     `xorm:"'email'" json:"email,omitempty"`
		Mobile           string     `xorm:"'mobile'" json:"mobile,omitempty"`
		IsEmailConfirmed bool       `xorm:"'is_email_confirmed'" json:"emailConfirmed"`
		Birthday         *time.Time `xorm:"'birthday'" json:"birthday,omitempty"` //生日
		LastLoginDate    *time.Time `xorm:"'last_login_date'" json:"lastLoginDate,omitempty"`
		CreateTime       *time.Time `xorm:"'create_time'" json:"createTime,omitempty"`
		UpdateTime       *time.Time `xorm:"'update_time'" json:"updateTime,omitempty"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
