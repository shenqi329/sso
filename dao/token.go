package dao

import (
	//"log"
	"sso/bean"
	"sso/mysql"
)

func InsertToken(token *bean.Token) error {

	engine := mysql.GetXormEngine()
	session := engine.NewSession()

	err := session.Begin()

	_, err = engine.Delete(bean.Token{DeviceId: token.DeviceId})
	if err != nil {
		session.Rollback()
		return err
	}

	_, err = engine.Insert(token)

	if err != nil {
		session.Rollback()
		return err
	}

	err = session.Commit()

	if err != nil {
		return err
	}

	return nil
}

func GetToken(token *bean.Token) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(token)

	return has, err
}

func RemoveToken(token *bean.Token) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Delete(&token)

	return count, err
}
