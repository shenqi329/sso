package dao

import (
	"errors"
	"log"
	"sso/bean"
	"sso/mysql"
)

var (
	ErrorDaoDBInnerFail = errors.New("数据库内部错误")
	ErrorDaoNotFound    = errors.New("数据未找到")
)

func GetUserByName(name string) (*bean.User, error) {

	engine := mysql.GetXormEngine()

	userBean := bean.User{Name: name}

	has, err := engine.Get(&userBean)

	if err != nil {
		log.Println(err.Error())
		return nil, ErrorDaoDBInnerFail
	}

	if has == false {
		return nil, ErrorDaoNotFound
	}

	return &userBean, nil

	// db := mysql.GetDB()

	// userBean := bean.User{}

	// db = db.Raw("SELECT * FROM `t_user`  WHERE (`t_user`.`user_username` = ?)", name).Find(&userBean)

	// if db.RecordNotFound() {
	// 	log.Println(db.Error.Error())
	// 	return nil, ErrorDaoNotFound
	// }
	// if err := db.Error; err != nil {
	// 	log.Println(db.Error.Error())
	// 	return nil, ErrorDaoDBInnerFail
	// }
	// log.Printf("name = %s,passpword = %s", userBean.Name, userBean.Password)
	// log.Print(userBean.ID)
	// log.Print(userBean.CreateTime)
	// return &userBean, nil
}

func InsertUser(user *bean.User) error {

	engine := mysql.GetXormEngine()

	engine.Insert(user)

	return nil

	db := mysql.GetDB()

	if err := db.Create(user).Error; err != nil {
		return ErrorDaoDBInnerFail
	}

	return nil
}

func GetUserById(id int64) (*bean.User, error) {

	engine := mysql.GetXormEngine()

	userBean := bean.User{ID: id}
	has, err := engine.Get(&userBean)

	if err != nil {
		log.Println(err.Error())
		return nil, ErrorDaoDBInnerFail
	}

	if has == false {
		return nil, ErrorDaoNotFound
	}

	return &userBean, nil
	// db := mysql.GetDB()

	// userBean := bean.User{}
	// db.Where("user_id = }", id).Find(&userBean)

	// return &userBean, nil
}
