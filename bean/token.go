package bean

import (
	"time"
)

type (
	Token struct {
		ID         int64     `xorm:"'token_id'" json:"id" bson:"_id,omitempty" gorm:"column:token_id"`
		UserId     int64     `xorm:"'token_user_id'" json:"userId" bson:"userId" form:"userId" gorm:"column:token_user_id"`
		Token      string    `xorm:"'token_token'" json:"token" bson:"token" form:"token" gorm:"column:token_token"`
		CreateTime time.Time `xorm:"'token_create_date'" json:"date" bson:"date" form:"date" gorm:"column:token_create_date"`
	}
)

func (u Token) TableName() string {
	return "t_token"
}
