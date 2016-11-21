package dao

import (
	"sso/bean"
	"sso/mysql"
)

func InsertEmail(token *bean.Email) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Insert(token)

	return count, err
}

func GetEmail(email *bean.Email) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(email)

	return has, err
}
