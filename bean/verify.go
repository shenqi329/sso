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
		ID          int64      `xorm:"'id'" json:"id,omitempty"`
		Type        int8       `xorm:"'type'" json:"type,omitempty"`
		VerifyId    string     `xorm:"'verify_id'" json:"verifyId,omitempty"`
		Code        string     `xorm:"'code'" json:"verifyCode,omitempty"`
		ExpiredTime *time.Time `xorm:"'expired_time'" json:"expiredTime"`
	}
)

func (v Verify) TableName() string {
	return "t_verify"
}
