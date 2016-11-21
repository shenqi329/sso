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
		ID               int64      `xorm:"'t_user_id'" json:"id,omitempty" gorm:"column:t_user_id;primary_key"`
		UserName         string     `xorm:"'t_user_username'" json:"username,omitempty" form:"username" gorm:"column:t_user_username"`
		Password         string     `xorm:"'t_user_password'" json:"password,omitempty" form:"password" gorm:"column:t_user_password"`
		Birthday         *time.Time `xorm:"'t_user_birthday'" json:"birthday,omitempty" gorm:"column:t_user_birthday"` //生日
		Name             *string    `xorm:"'t_user_name'" json:"name,omitempty" form:"name" gorm:"column:t_user_name"` //姓名
		Icon             *string    `xorm:"'t_user_icon'" json:"icon,omitempty" form:"icon" gorm:"column:t_user_icon"` //头像图片url
		Sex              *string    `xorm:"'t_user_sex'" json:"sex,omitempty" form:"sex" gorm:"column:t_user_sex"`     //性别
		NickName         *string    `xorm:"'t_user_nickname'" json:"nickname,omitempty" form:"nickname" gorm:"column:t_user_nickname"`
		LastLoginDate    *time.Time `xorm:"'t_user_last_login_date'" json:"lastLoginDate,omitempty" form:"lastLoginDate" gorm:"column:t_user_last_login_date"`
		Email            string     `xorm:"'t_user_email'" json:"email,omitempty" form:"email" gorm:"column:t_user_email"`
		Mobile           *string    `xorm:"'t_user_mobile'" json:"mobile,omitempty" gorm:"column:t_user_mobile"`
		CreateTime       *time.Time `xorm:"'t_user_create_time'" json:"createTime" bson:"createTime" form:"createTime" gorm:"column:t_user_create_date"`
		UpdateTime       *time.Time `xorm:"'t_user_update_time'" json:"updateTime" bson:"updateTime" form:"updateTime" gorm:"column:t_user_update_time"`
		IsEmailConfirmed bool       `xorm:"'t_user_is_email_confirmed'" json:"emailConfirmed" bson:"emailConfirmed,omitempty" gorm:"column:t_user_is_email_confirmed"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
