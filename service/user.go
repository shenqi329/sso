package service

import (
	"errors"
	"log"
	"sso/bean"
	"sso/dao"
	ssoerror "sso/error"
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
		return nil, ssoerror.ErrorIllegalParams
		return nil, ErrorParams
	}

	tokenBean := &bean.Token{Token: token}
	userBean := &bean.User{}

	has, err := dao.GetToken(tokenBean)

	if has {
		has, err = dao.GetUser(userBean)
	}

	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
		return nil, ErrorDatabaseOperation
	}

	if !has {
		return nil, ssoerror.ErrorNotFound
		return nil, ErrorNotFound
	}

	return userBean, nil
}

func UserLogin(user *bean.User, token *bean.Token) (*bean.User, *bean.Token, error) {

	has, err := dao.GetUser(user)

	if err != nil {
		return nil, nil, ssoerror.ErrorInternalServerError
		return nil, nil, err
	}
	if has == false {
		return nil, nil, ssoerror.ErrorNotFound
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
		log.Println(err.Error())
		return nil, nil, ssoerror.ErrorInternalServerError
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
		return nil, nil, ssoerror.ErrorInternalServerError
		return ErrorDatabaseOperation
	}
	if count <= 0 {
		return nil, nil, ssoerror.ErrorNotFound
		return ErrorNotFound
	}

	return nil
}

func UserRegister(user *bean.User, email *bean.Email) (*bean.User, error) {

	if len(user.UserName) == 0 || len(user.Password) == 0 {
		return nil, ssoerror.ErrorIllegalParams
		return nil, ErrorParams
	}

	has, err := dao.GetUser(&bean.User{UserName: user.UserName})
	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
		return nil, ErrorDatabaseOperation
	}
	if has {
		return nil, ssoerror.ErrorResourceExist
		return nil, ErrorResourceExist
	}

	has, err = dao.GetEmail(&bean.Email{Email: email.Email, UserName: email.UserName, Code: email.Code})
	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
		return nil, ErrorDatabaseOperation
	}

	if !has {
		log.Println("验证码错误")
		return nil, ssoerror.ErrorRegisterErrorCode
		return nil, ErrorDatabaseOperation
	}

	timeNow := time.Now()
	user.CreateTime = &timeNow
	user.UpdateTime = &timeNow
	user.IsEmailConfirmed = true

	_, err = dao.InsertUser(user)

	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
		return nil, ErrorDatabaseOperation
	}

	return user, nil
}
