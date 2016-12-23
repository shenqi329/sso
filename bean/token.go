package bean

import (
	"time"
)

type (
	Token struct {
		ID          int64      `xorm:"'id'" json:"id"`
		UserId      int64      `xorm:"'user_id'" json:"userId"`           //用户id
		Token       string     `xorm:"'token'" json:"token"`              //凭证
		DeviceId    string     `xorm:"'device_id'" json:"deviceId"`       //app的登录时的传入的uuid
		AppId       string     `xorm:"'app_id'" json:"appId"`             //分配给app的固定值,一个app对应一个唯一的id ,如easynote的appId为 89897
		Platform    string     `xorm:"'platform'" json:"platform"`        //平台:ios, android,windows,mac
		CreateTime  *time.Time `xorm:"'create_time'" json:"createTime"`   //创建时间
		ExpiredTime *time.Time `xorm:"'expired_time'" json:"expiredTime"` //过期时间
	}
)

func (u Token) TableName() string {
	return "t_token"
}
