package bean

import (
	"time"
)

type (
	Token struct {
		ID          int64      `xorm:"'t_token_id'" json:"id"`
		UserId      int64      `xorm:"'t_token_user_id'" json:"userId"`           //用户id
		Token       string     `xorm:"'t_token_token'" json:"token"`              //凭证
		DeviceId    string     `xorm:"'t_token_device_id'" json:"deviceId"`       //app的登录时的传入的uuid
		AppId       string     `xorm:"'t_token_app_id'" json:"appId"`             //分配给app的固定值,一个app对应一个唯一的id ,如easynote的appId为 89897
		Platform    string     `xorm:"'t_token_platform'" json:"platform"`        //平台:ios, android,windows,mac
		CreateTime  *time.Time `xorm:"'t_token_create_time'" json:"createTime"`   //创建时间
		ExpiredTime *time.Time `xorm:"'t_token_expired_time'" json:"expiredTime"` //过期时间
	}
)

func (u Token) TableName() string {
	return "t_token"
}
