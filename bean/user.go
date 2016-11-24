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
		ID               int64      `xorm:"'t_user_id'" json:"id,omitempty"`
		UserName         string     `xorm:"'t_user_username'" json:"username,omitempty" form:"username"`
		Password         string     `xorm:"'t_user_password'" json:"password,omitempty" form:"password"`
		Name             string     `xorm:"'t_user_name'" json:"name,omitempty" form:"name"` //姓名
		Icon             string     `xorm:"'t_user_icon'" json:"icon,omitempty" form:"icon"` //头像图片url
		Sex              string     `xorm:"'t_user_sex'" json:"sex,omitempty" form:"sex"`    //性别
		NickName         string     `xorm:"'t_user_nickname'" json:"nickname,omitempty" form:"nickname"`
		Email            string     `xorm:"'t_user_email'" json:"email,omitempty" form:"email"`
		Mobile           string     `xorm:"'t_user_mobile'" json:"mobile,omitempty"`
		IsEmailConfirmed bool       `xorm:"'t_user_is_email_confirmed'" json:"emailConfirmed"`
		Birthday         *time.Time `xorm:"'t_user_birthday'" json:"birthday,omitempty" form:"birthday"` //生日
		LastLoginDate    *time.Time `xorm:"'t_user_last_login_date'" json:"lastLoginDate,omitempty" form:"lastLoginDate"`
		CreateTime       *time.Time `xorm:"'t_user_create_time'" json:"createTime,omitempty" form:"createTime"`
		UpdateTime       *time.Time `xorm:"'t_user_update_time'" json:"updateTime,omitempty" form:"updateTime"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
