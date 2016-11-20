package bean

import (
	"time"
)

type (
	Token struct {
		ID          int64      `xorm:"'t_token_id'" json:"id" bson:"_id,omitempty" gorm:"column:t_token_id"`
		UserId      int64      `xorm:"'t_token_user_id'" json:"userId" bson:"userId" form:"userId" gorm:"column:t_token_user_id"`                             //用户id
		Token       string     `xorm:"'t_token_token'" json:"token" bson:"token" form:"token" gorm:"column:t_token_token"`                                    //凭证
		DeviceId    string     `xorm:"'t_token_device_id'" json:"deviceId" bson:"deviceId" form:"deviceId" gorm:"column:t_token_device_id"`                   //app的登录时的传入的uuid
		AppId       string     `xorm:"'t_token_app_id'" json:"appId" bson:"appId" form:"appId" gorm:"column:t_token_app_id"`                                  //分配给app的固定值,一个app对应一个唯一的id ,如easynote的appId为 89897
		Platform    string     `xorm:"'t_token_platform'" json:"platform" bson:"platform" form:"platform" gorm:"column:t_token_platform"`                     //平台:ios, android,windows,mac
		CreateTime  *time.Time `xorm:"'t_token_create_time'" json:"createTime" bson:"createTime" form:"createTime" gorm:"column:t_token_create_time"`         //创建时间
		ExpiredTime *time.Time `xorm:"'t_token_expired_time'" json:"expairedTime" bson:"expairedTime" form:"expairedTime" gorm:"column:t_token_expired_time"` //过期时间
	}
)

func (u Token) TableName() string {
	return "t_token"
}
