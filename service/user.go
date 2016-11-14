package service

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sso/bean"
	"sso/mongodb"
)

var (
	ErrorUserParams      = errors.New("参数错误")
	ErrorUserNoExist     = errors.New("用户不存在")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserDBTableFail = errors.New("获取数据库表失败")
	ErrorUserDBFindError = errors.New("数据查询错误")
	ErrorUserNotFound    = errors.New("数据未找到")
)

func GetUser(name string) (*bean.User, error) {

	session := mongodb.GetSession()

	c := session.DB("db_sso").C("t_user")
	if c == nil {
		return nil, ErrorUserDBTableFail
	}

	user := bean.User{}
	err := c.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		fmt.Println("getuser,err" + err.Error())
		if err == mgo.ErrNotFound {
			return nil, ErrorUserNotFound //说明没有查询到，用户不存在
		}
		return nil, ErrorUserDBFindError //查询发生错误
	}
	fmt.Println("name:" + user.Name)
	return &user, nil //查询
}

func UserLogin(user *bean.User) (*bean.User, error) {

	if len(user.Name) == 0 || len(user.Password) == 0 {
		return nil, ErrorUserParams
	}

	result, err := GetUser(user.Name)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrorUserNoExist
	}

	return result, nil
}

func UserRegister(user *bean.User) (*bean.User, error) {

	if len(user.Name) == 0 || len(user.Password) == 0 {
		return nil, ErrorUserParams
	}

	session := mongodb.GetSession()

	c := session.DB("db_sso").C("t_user")
	if c == nil {
		return nil, ErrorUserDBTableFail
	}

	result, err := GetUser(user.Name)
	if err != nil && err != ErrorUserNotFound {
		return nil, err
	}
	if result != nil {
		return nil, ErrorUserExist
	}

	err = c.Insert(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
