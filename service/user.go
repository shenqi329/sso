package service

import (
	"log"
	"sso/bean"
	"sso/dao"
	ssoerror "sso/error"
	"sso/uuid"
	"strings"
	"time"
)

func UserInfoByToken(token string) (*bean.User, error) {

	if err := CheckToken(token); err != nil {
		return nil, err
	}

	tokenBean := &bean.Token{Token: token}

	has, err := dao.GetToken(tokenBean)
	userBean := &bean.User{ID: tokenBean.UserId}

	if has {
		has, err = dao.GetUser(userBean)
	}

	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
	}

	if !has {
		return nil, ssoerror.ErrorTokenInvalidated
	}

	return userBean, nil
}

func UserLogin(user *bean.User, token *bean.Token) (*bean.User, *bean.Token, error) {

	has, err := dao.GetUser(user)

	if err != nil {
		return nil, nil, ssoerror.ErrorInternalServerError
	}
	if has == false {
		return nil, nil, ssoerror.ErrorNotFound
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
	}
	lastLoginDate := time.Now()
	dao.UpdateUser(&bean.User{LastLoginDate: &lastLoginDate}, &bean.User{ID: user.ID})
	return user, tokenBean, nil
}

func UserLogout(token string) error {

	if err := CheckToken(token); err != nil {
		return err
	}

	count, err := dao.RemoveToken(&bean.Token{Token: token})

	if err != nil {
		return ssoerror.ErrorInternalServerError
	}
	if count <= 0 {
		return ssoerror.ErrorNotFound
	}

	return nil
}

func UserChangePassword(token string, userName string, originalPassword string, newPassword string) error {

	if userName == "" ||
		originalPassword == "" {
		return ssoerror.ErrorIllegalParams
	}

	if err := CheckPassword(newPassword); err != nil {
		return ssoerror.ErrorPasswordFormatError
	}

	if strings.EqualFold(originalPassword, newPassword) {
		return ssoerror.ErrorSameOriginalNewPassword
	}

	userBean, err := UserInfoByToken(token)
	if err != nil {
		return err
	}
	if !strings.EqualFold(userBean.UserName, userName) {
		return ssoerror.ErrorIllegalParams
	}

	count, err := dao.UpdateUser(&bean.User{Password: newPassword}, &bean.User{UserName: userName, Password: originalPassword})
	if err != nil {
		return ssoerror.ErrorInternalServerError
	}
	if count <= 0 {
		return ssoerror.ErrorUserNameOrPasswordFail
	}
	return nil
}

func UserRegister(user *bean.User, email *bean.Email) (*bean.User, error) {

	if err := CheckUserName(user.UserName); err != nil {
		return nil, err
	}
	if err := CheckPassword(user.Password); err != nil {
		return nil, err
	}
	if err := CheckEmail(email.Email); err != nil {
		return nil, err
	}

	has, err := dao.GetUser(&bean.User{UserName: user.UserName})
	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
	}
	if has {
		return nil, ssoerror.ErrorRegisterUserExist
	}

	has, err = dao.GetEmail(&bean.Email{Email: email.Email, UserName: email.UserName, Code: email.Code})
	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
	}

	if !has {
		return nil, ssoerror.ErrorRegisterErrorCode
	}

	timeNow := time.Now()
	user.CreateTime = &timeNow
	user.UpdateTime = &timeNow
	user.IsEmailConfirmed = true

	_, err = dao.InsertUser(user)

	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
	}

	return user, nil
}
