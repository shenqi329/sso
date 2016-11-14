package bean

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Token struct {
		ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		UserId     string        `json:"userId" bson:"userId" form:"userId"`
		Token      string        `json:"token" bson:"token" form:"token"`
		CreateTime time.Time     `json:"date" bson:"date" form:"date"`
	}
)
