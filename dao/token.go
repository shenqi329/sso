package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sso/bean"
	"sso/mongodb"
)

func InsertToken(token *bean.Token) error {

	session := mongodb.GetSession()
	c := session.DB("db_sso").C("t_token")

	if c == nil {
		return ErrorDaoDBInnerFail
	}

	err := c.Insert(token)

	if err != nil {
		return ErrorDaoDBInnerFail
	}
	return nil
}

func GetTokenByToken(token string) (*bean.Token, error) {

	session := mongodb.GetSession()

	c := session.DB("db_sso").C("t_user")
	if c == nil {
		return nil, ErrorDaoDBInnerFail
	}

	tokenBean := bean.Token{}
	err := c.Find(bson.M{"token": token}).One(&tokenBean)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrorDaoNotFound
		}
		return nil, ErrorDaoDBInnerFail
	}

	return &tokenBean, nil
}

func RemoveTokenByToken(token string) error {
	session := mongodb.GetSession()

	c := session.DB("db_sso").C("t_user")
	if c == nil {
		return ErrorDaoDBInnerFail
	}

	err := c.Remove(bson.D{{"token", token}})
	if err != nil {
		if err == mgo.ErrNotFound {
			return ErrorDaoNotFound
		}
		return ErrorDaoDBInnerFail
	}
	return nil
}
