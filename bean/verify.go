package bean

import (
	"time"
)

const (
	VerifyTypeRegisterEmail  = 1
	VerifyTypeRegisterMobile = 2
	VerifyTypeChangeEmail    = 3
	VerifyTypeChangeMobile   = 4
)

type (
	Verify struct {
		ID          int64      `xorm:"'t_verify_id'" json:"id,omitempty"`
		Type        int8       `xorm:"'t_verify_verify_type'" json:"type,omitempty"`
		VerifyId    string     `xorm:"'t_verify_verify_id'" json:"verifyId,omitempty"`
		VerifyCode  string     `xorm:"'t_verify_code'" json:"verifyCode,omitempty"`
		ExpiredTime *time.Time `xorm:"'t_verify_expired_time'" json:"expiredTime"`
	}
)

func (v Verify) TableName() string {
	return "t_verify"
}
