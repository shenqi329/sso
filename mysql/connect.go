package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		db = Connect()
	}
	return db
}

func Connect() *gorm.DB {

	db, err := gorm.Open("mysql", "user_connect:user_connect@tcp(172.17.0.2:3306)/db_sso")

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return db
}
