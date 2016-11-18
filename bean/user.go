package bean

import (
	"time"
)

const (
	UserGenderMale   = "m"
	UserGenderFemale = "f"
)

type (
	User struct {
		ID       int64  `xorm:"'t_user_id'" json:"id" bson:"_id,omitempty" gorm:"column:t_user_id;primary_key"`
		Name     string `xorm:"'t_user_username'" json:"name" bson:"name" form:"name" gorm:"column:t_user_username"`
		Password string `xorm:"'t_user_password'" json:"password" bson:"password" form:"password" gorm:"column:t_user_password"`

		CreateTime time.Time `xorm:"'t_user_create_time'" json:"createTime" bson:"createTime" form:"createTime" gorm:"column:t_user_create_date"`
		UpdateTime time.Time `xorm:"'t_user_update_time'" json:"updateTime" bson:"updateTime" form:"updateTime" gorm:"column:t_user_update_time"`
	}

	UserExist struct {
		Name  string `json:"name" xml:"name" form:"name"`
		Exist bool   `json:"exist" xml:"exist" form:"exist"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
