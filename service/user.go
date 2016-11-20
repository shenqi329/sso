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
	ErrorResourceExist     = errors.New("resource already exist")
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

	createTime := time.Now()
	expiredTime := createTime.Add(2 * time.Hour)
	tokenBean := &bean.Token{
		AppId:       token.AppId,
		DeviceId:    token.DeviceId,
		Platform:    token.Platform,
		UserId:      user.ID,
		Token:       uuid.Rand().Hex(),
		ExpiredTime: &expiredTime,
		CreateTime:  &createTime,
	}

	err = dao.InsertToken(tokenBean)
	if err != nil {
		return nil, nil, ErrorDatabaseOperation
	}

	lastLoginDate := time.Now()
	dao.UpdateUser(&bean.User{ID: user.ID, LastLoginDate: &lastLoginDate})

	return user, tokenBean, nil
}

func UserLogout(token string) error {

	if len(token) <= 0 {
		return ErrorParams
	}

	count, err := dao.RemoveToken(&bean.Token{Token: token})

	if err != nil {
		return ErrorDatabaseOperation
	}
	if count <= 0 {
		return ErrorNotFound
	}

	return nil
}

func UserRegister(user *bean.User) (*bean.User, error) {

	if len(user.UserName) == 0 || len(user.Password) == 0 {
		return nil, ErrorParams
	}

	has, err := dao.GetUser(&bean.User{UserName: user.UserName})
	if err != nil {
		return nil, ErrorDatabaseOperation
	}

	if has {

		return nil, ErrorResourceExist
	}

	timeNow := time.Now()

	user.CreateTime = &timeNow
	user.UpdateTime = &timeNow

	_, err = dao.InsertUser(user)

	if err != nil {
		return nil, ErrorDatabaseOperation
	}

	return user, nil
}
