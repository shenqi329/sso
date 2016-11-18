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
		ID       int64  `xorm:"'user_id'" json:"id" bson:"_id,omitempty" "gorm:"column:user_id;primary_key"`
		Name     string `xorm:"'user_username'" json:"name" bson:"name" form:"name" gorm:"column:user_username"`
		Password string `xorm:"'user_password'" json:"password" bson:"password" form:"password" gorm:"column:user_password"`

		CreateTime time.Time `xorm:"'user_create_date'" json:"createTime" bson:"createTime" form:"createTime" gorm:"-;column:user_create_date"`
		UpdateTime time.Time `xorm:"'user_update_date'" json:"updateTime" bson:"updateTime" form:"updateTime" gorm:"-;column:user_update_date"`
	}

	UserExist struct {
		Name  string `json:"name" xml:"name" form:"name"`
		Exist bool   `json:"exist" xml:"exist" form:"exist"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
