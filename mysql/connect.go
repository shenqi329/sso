package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var engine *xorm.Engine

func GetDB() *gorm.DB {
	if db == nil {
		db = connect()
	}
	return db
}

func GetXormEngine() *xorm.Engine {
	if engine == nil {
		eng, err := xorm.NewEngine("mysql", "user_connect:user_connect@tcp(172.17.0.2:3306)/db_sso?charset=utf8")
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		engine = eng
	}
	return engine
}

func connect() *gorm.DB {

	db, err := gorm.Open("mysql", "user_connect:user_connect@tcp(172.17.0.2:3306)/db_sso?charset=utf8")

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	db.LogMode(true)

	return db
}
