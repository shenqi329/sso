package dao

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sso/bean"
	"sso/mysql"
)

var (
	ErrorDaoDBInnerFail = errors.New("数据库内部错误")
	ErrorDaoNotFound    = errors.New("数据未找到")
)

func GetUserByName(name string) (*bean.User, error) {

	db := mysql.GetDB()

	userBean := bean.User{}
	db.Where("name = ", name).Find(&userBean)

	return &userBean, nil

	// session := mongodb.GetSession()

	// c := session.DB("db_sso").C("t_user")
	// if c == nil {
	// 	return nil, ErrorDaoDBInnerFail
	// }

	// user := bean.User{}
	// err := c.Find(bson.M{"name": name}).One(&user)
	// if err != nil {
	// 	if err == mgo.ErrNotFound {
	// 		return nil, ErrorDaoNotFound //说明没有查询到，用户不存在
	// 	}
	// 	return nil, ErrorDaoDBInnerFail //查询发生错误
	// }
	// return &user, nil //查询
}

func InsertUser(user *bean.User) error {

	db := mysql.GetDB()

	db.Create(user)

	return nil
	// session := mongodb.GetSession()

	// c := session.DB("db_sso").C("t_user")
	// if c == nil {
	// 	return ErrorDaoDBInnerFail
	// }
	// err := c.Insert(user)

	// if err != nil {
	// 	return ErrorDaoDBInnerFail //查询发生错误
	// }
	// return nil
}

func GetUserById(id string) (*bean.User, error) {

	db := mysql.GetDB()

	userBean := bean.User{}
	db.Where("id = ", id).Find(out, &userBean)

	// session := mongodb.GetSession()

	// c := session.DB("db_sso").C("t_user")
	// if c == nil {
	// 	return nil, ErrorDaoDBInnerFail
	// }

	// user := bean.User{}
	// err := c.FindId(bson.ObjectIdHex(id)).One(&user)
	// if err != nil {
	// 	if err == mgo.ErrNotFound {
	// 		return nil, ErrorDaoNotFound //说明没有查询到，用户不存在
	// 	}
	// 	return nil, ErrorDaoDBInnerFail //查询发生错误
	// }
	// return &user, nil //查询
}
