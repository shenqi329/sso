package service

import (
	"errors"
	//"log"
	"sso/bean"
	"sso/dao"
	"sso/uuid"
	"time"
)

var (
	ErrorParams            = errors.New("parameter error")
	ErrorResourceExists    = errors.New("resource already exists")
	ErrorNotFound          = errors.New("resource doesn't exist")
	ErrorDatabaseOperation = errors.New("database operation")
)

func UserInfo(token string) (*bean.User, error) {

	if len(token) <= 0 {
		return nil, ErrorParams
	}

	tokenBean := &bean.Token{Token: token}
	userBean := &bean.User{}

	has, err := dao.GetToken(tokenBean)

	if has {
		has, err = dao.GetUser(userBean)
	}

	if err != nil {
		return nil, ErrorDatabaseOperation
	}

	if !has {
		return nil, ErrorNotFound
	}

	return userBean, nil
}

func UserLogin(user *bean.User, token *bean.Token) (*bean.User, *bean.Token, error) {

	has, err := dao.GetUser(user)

	if err != nil {
		return nil, nil, err
	}
	if has == false {
		return nil, nil, ErrorNotFound
	}

	tokenBean := &bean.Token{
		AppId:        token.AppId,
		DeviceId:     token.DeviceId,
		Platform:     token.Platform,
		UserId:       user.ID,
		Token:        uuid.Rand().Hex(),
		ExpairedTime: time.Now().Add(2 * time.Hour),
		CreateTime:   time.Now(),
	}

	err = dao.InsertToken(tokenBean)
	if err != nil {
		return nil, nil, ErrorDatabaseOperation
	}

	dao.UpdateUser(&bean.User{ID: user.ID, LastLoginDate: time.Now()})

	return user, tokenBean, nil
}

func UserLogout(token string) error {

	if len(token) <= 0 {
		return ErrorParams
	}

	_, err := dao.RemoveToken(&bean.Token{Token: token})

	if err != nil {
		return ErrorDatabaseOperation
	}

	return nil
}

func UserRegister(user *bean.User) (*bean.User, error) {

	if len(user.UserName) == 0 || len(user.Password) == 0 {
		return nil, ErrorParams
	}

	has, err := dao.GetUser(user)

	if err != nil {
		return nil, ErrorDatabaseOperation
	}

	if has {

		return nil, ErrorResourceExists
	}

	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()

	_, err = dao.InsertUser(user)

	if err != nil {
		return nil, ErrorDatabaseOperation
	}

	return user, nil
}
