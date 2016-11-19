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
		ID            int64     `xorm:"'t_user_id'" json:"id" gorm:"column:t_user_id;primary_key"`
		UserName      string    `xorm:"'t_user_username'" json:"name" form:"name" gorm:"column:t_user_username"`
		Password      string    `xorm:"'t_user_password'" json:"password" form:"password" gorm:"column:t_user_password"`
		Birthday      time.Time `xorm:"'t_user_birthday'" json:"birthday" gorm:"column:t_user_birthday"` //生日
		Name          *string   `xorm:"'t_user_name'" json:"name" gorm:"column:t_user_name"`             //姓名
		Icon          *string   `xorm:"'t_user_icon'" json:"icon" gorm:"column:t_user_icon"`             //头像图片url
		Sex           *string   `xorm:"'t_user_sex'" json:"sex" gorm:"column:t_user_sex"`                //性别
		NickName      *string   `xorm:"'t_user_nickname'" json:"nickname" gorm:"column:t_user_nickname"`
		LastLoginDate time.Time `xorm:"'t_user_last_login_date'" json:"last_login_date" gorm:"column:t_user_last_login_date"`
		Email         *string   `xorm:"'t_user_email'" json:"email" gorm:"column:t_user_email"`
		Mobile        *string   `xorm:"'t_user_mobile'" json:"mobile" gorm:"column:t_user_mobile"`
		CreateTime    time.Time `xorm:"'t_user_create_time'" json:"createTime" bson:"createTime" form:"createTime" gorm:"column:t_user_create_date"`
		UpdateTime    time.Time `xorm:"'t_user_update_time'" json:"updateTime" bson:"updateTime" form:"updateTime" gorm:"column:t_user_update_time"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
