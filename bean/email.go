package bean

import (
	"time"
)

type (
	Email struct {
		ID          int64      `xorm:"'t_email_verify_id'" json:"id,omitempty" form:"id"`
		UserName    string     `xorm:"'t_email_verify_username'" json:"username,omitempty" form:"username"`
		Code        string     `xorm:"'t_email_verify_code'" json:"code,omitempty" form:"code"`
		ExpiredTime *time.Time `xorm:"'t_email_verify_expired_time'" json:"expiredTime" form:"expiredTime"`
	}
)

func (e Email) TableName() string {
	return "t_email_verify"
}
