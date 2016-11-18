package bean

import (
	"time"
)

type (
	Token struct {
		ID           int64     `xorm:"'t_token_id'" json:"id" bson:"_id,omitempty" gorm:"column:t_token_id"`
		UserId       int64     `xorm:"'t_token_user_id'" json:"userId" bson:"userId" form:"userId" gorm:"column:t_token_user_id"`
		Token        string    `xorm:"'t_token_token'" json:"token" bson:"token" form:"token" gorm:"column:t_token_token"`
		CreateTime   time.Time `xorm:"'t_token_create_time'" json:"date" bson:"date" form:"date" gorm:"column:t_token_create_time"`
		ExpairedTime time.Time `xorm:"'t_token_expaired_time'" json:"date" bson:"date" form:"date" gorm:"column:t_token_expaired_time"`
	}
)

func (u Token) TableName() string {
	return "t_token"
}
