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
		ID       int64  `json:"id" bson:"_id,omitempty" "gorm:"column:user_id"`
		Name     string `json:"name" bson:"name" form:"name" gorm:"column:user_username"`
		Password string `json:"password" bson:"password" form:"password" gorm:"column:user_password"`

		CreateTime time.Time `json:"createTime" bson:"createTime" form:"createTime" gorm:"column:user_create_date"`
		UpdateTime time.Time `json:"createTime" bson:"createTime" form:"createTime" gorm:"column:user_update_date"`
	}

	UserExist struct {
		Name  string `json:"name" xml:"name" form:"name"`
		Exist bool   `json:"exist" xml:"exist" form:"exist"`
	}
)

func (u User) TableName() string {
	return "t_user"
}
