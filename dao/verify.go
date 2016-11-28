package dao

import (
	"sso/bean"
	"sso/mysql"
)

func InsertVerify(verifys ...*bean.Verify) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Insert(verifys)

	return count, err
}

func GetVerify(verify *bean.Verify) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(verify)

	return has, err
}
