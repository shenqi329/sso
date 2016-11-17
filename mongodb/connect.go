package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		session = connect()
	}
	return session
}

func connect() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		fmt.Println("mongodb:连接失败")
		return nil
	}
	fmt.Println("mongodb:连接成功")

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("db_sso").C("t_user")

	index := mgo.Index{
		Key:         []string{"CreateTime"},
		ExpireAfter: 1 * time.Minute,
	}
	c.EnsureIndex(index)

	return session
}
