package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sso/bean"
	"sso/mysql"
)

func InsertToken(token *bean.Token) error {

	db := mysql.GetDB()
	db.Save(token)

	return nil

	// session := mongodb.GetSession()
	// c := session.DB("db_sso").C("t_token")

	// if c == nil {
	// 	return ErrorDaoDBInnerFail
	// }

	// err := c.Insert(token)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	return ErrorDaoDBInnerFail
	// }
	// return nil
}

func GetTokenByToken(token string) (*bean.Token, error) {

	db := mysql.GetDB()

	tokenBean := bean.Token{}

	db.Where("token_token = ?", token).First(&tokenBean)

	return &tokenBean, nil
	// session := mongodb.GetSession()

	// c := session.DB("db_sso").C("t_token")
	// if c == nil {
	// 	log.Println(ErrorDaoDBInnerFail.Error())
	// 	return nil, ErrorDaoDBInnerFail
	// }

	// tokenBean := bean.Token{}
	// err := c.Find(bson.M{"token": token}).One(&tokenBean)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	if err == mgo.ErrNotFound {
	// 		return nil, ErrorDaoNotFound
	// 	}
	// 	return nil, ErrorDaoDBInnerFail
	// }

	// return &tokenBean, nil
}

func RemoveTokenByToken(token string) error {

	db := mysql.GetDB()

	db.Where("token_token = ", token).Delete(bean.Token{})

	// session := mongodb.GetSession()

	// c := session.DB("db_sso").C("t_token")
	// if c == nil {
	// 	return ErrorDaoDBInnerFail
	// }

	// err := c.Remove(bson.D{{"token", token}})
	// if err != nil {
	// 	log.Println(err.Error())
	// 	if err == mgo.ErrNotFound {
	// 		return ErrorDaoNotFound
	// 	}
	// 	return ErrorDaoDBInnerFail
	// }
	// return nil
}
