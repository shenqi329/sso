package dao

import (
	//"errors"
	//"log"
	"sso/bean"
	"sso/mysql"
)

func GetUser(user *bean.User) (bool, error) {

	engine := mysql.GetXormEngine()

	has, err := engine.Get(user)

	return has, err
}

func UpdateUser(user *bean.User, condiUsers ...interface{}) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Update(user, condiUsers...)

	return count, err
}

func InsertUser(user *bean.User) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Insert(user)

	return count, err
}
