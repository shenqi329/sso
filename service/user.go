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
		log.Println(err.Error())
		return nil, ssoerror.ErrorInternalServerError
	}

	if !has {
		return nil, ssoerror.ErrorTokenInvalidated
	}

	return userBean, nil
}

func UserLogin(user *bean.User, token *bean.Token) (*bean.User, *bean.Token, error) {

	if len(user.UserName) <= 0 ||
		len(user.Password) <= 0 {
		log.Println("获取用户信息失败")
		return nil, nil, ssoerror.ErrorIllegalParams
	}

	if len(token.DeviceId) <= 0 ||
		len(token.AppId) <= 0 ||
		len(token.Platform) <= 0 {
		log.Println("获取授权信息失败")
		return nil, nil, ssoerror.ErrorIllegalParams
	}

	if err := CheckDeviceId(token.DeviceId); err != nil {
		return nil, nil, err
	}

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

func UserUpdate(token string, user *bean.User) error {

	userBean, err := UserInfoByToken(token)
	if err != nil {
		return err
	}
	log.Println(bean.StructToJsonString(userBean))

	//目前只有这些属性能够被这个接口修改
	updateUser := &bean.User{
		Birthday: user.Birthday,
		Name:     user.Name,
		Icon:     user.Icon,
		Sex:      user.Sex,
		NickName: user.NickName,
	}

	_, err = dao.UpdateUser(updateUser, &bean.User{ID: userBean.ID, UserName: userBean.UserName})
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}

	return nil
}

func UserChangePassword(token string, originalPassword string, newPassword string) error {

	if originalPassword == "" {
		return ssoerror.ErrorIllegalParams
	}

	if err := CheckPassword(newPassword); err != nil {
		return ssoerror.ErrorPasswordFormatError
	}

	userBean, err := UserInfoByToken(token)
	if err != nil {
		return err
	}

	if !strings.EqualFold(userBean.Password, originalPassword) {
		return ssoerror.ErrorPasswordWrong
	}

	if strings.EqualFold(userBean.Password, newPassword) {
		return ssoerror.ErrorSameOriginalNewPassword
	}

	count, err := dao.UpdateUser(&bean.User{Password: newPassword}, &bean.User{UserName: userBean.UserName, Password: originalPassword})
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
		log.Println(err.Error())
		return nil, ssoerror.ErrorInternalServerError
	}

	if !has {
		log.Println("没有查询到数据")
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
