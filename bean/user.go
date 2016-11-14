package bean

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	UserGenderMale   = "m"
	UserGenderFemale = "f"
)

type (
	User struct {
		ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name" form:"name"`
		Password   string        `json:"password" bson:"password" form:"password"`
		NickName   string        `json:"nickname" bson:"nickname" form:"nickname"`
		CreateTime time.Time     `json:"createTime" bson:"createTime" form:"createTime"`
	}
)

type (
	UserExist struct {
		Name  string `json:"name" xml:"name" form:"name"`
		Exist bool   `json:"exist" xml:"exist" form:"exist"`
	}
)
