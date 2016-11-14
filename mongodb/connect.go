package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		session = Connect()
	}
	return session
}

func Connect() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		fmt.Println("mongodb:连接失败")
		//panic(err)
		return nil
	}
	fmt.Println("mongodb:连接成功")

	session.SetMode(mgo.Monotonic, true)

	return session
}
