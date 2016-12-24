package service

import (
	"log"
	"sso/bean"
	"sso/dao"
	ssoerror "sso/error"
	"sso/request"
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

func UserLogin(login *request.Login) (*bean.User, *bean.Token, error) {

	if len(login.UserName) <= 0 ||
		len(login.Password) <= 0 {
		log.Println("获取用户信息失败")
		return nil, nil, ssoerror.ErrorIllegalParams
	}

	if len(login.DeviceId) <= 0 ||
		len(login.AppId) <= 0 ||
		len(login.Platform) <= 0 {
		log.Println("获取授权信息失败")
		return nil, nil, ssoerror.ErrorIllegalParams
	}

	if err := CheckDeviceId(login.DeviceId); err != nil {
		return nil, nil, err
	}

	user := &bean.User{UserName: login.UserName, Password: login.Password}
	token := &bean.Token{DeviceId: login.DeviceId, AppId: login.AppId, Platform: login.Platform}

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
		ExpiredTime: expiredTime,
		CreateTime:  createTime,
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

func UserUpdate(token string, update *request.Update) error {

	log.Println(bean.StructToJsonString(update))

	userBean, err := UserInfoByToken(token)
	if err != nil {
		return err
	}

	var birthday *time.Time
	if update.Birthday != nil {
		time := time.Unix(*update.Birthday, 0)
		birthday = &time
	}

	updateUser := &bean.User{
		Birthday: birthday,
		Name:     update.Name,
		Icon:     update.Icon,
		Sex:      update.Sex,
		NickName: update.NickName,
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

func ChangeEmail(token string, newEmail string, verifyCode string) error {
	if err := CheckToken(token); err != nil {
		return err
	}
	if err := CheckEmail(newEmail); err != nil {
		return err
	}
	if err := CheckVerifyCode(verifyCode); err != nil {
		return err
	}

	has, err := dao.GetVerify(&bean.Verify{Type: bean.VerifyTypeChangeEmail, VerifyId: newEmail, Code: verifyCode})
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}
	if !has {
		return ssoerror.ErrorRegisterErrorVerifyCode
	}

	userBean, err := UserInfoByToken(token)
	if err != nil {
		return err
	}

	updateUser := &bean.User{
		Email: newEmail,
	}

	_, err = dao.UpdateUser(updateUser, &bean.User{ID: userBean.ID, UserName: userBean.UserName})
	if err != nil {
		log.Println(err.Error())
		return ssoerror.ErrorInternalServerError
	}

	return nil
}

func UserRegister(register *request.Register) (*bean.User, error) {

	if err := CheckUserName(register.UserName); err != nil {
		return nil, err
	}
	if err := CheckPassword(register.Password); err != nil {
		return nil, err
	}
	if err := CheckEmail(register.Email); err != nil {
		return nil, err
	}

	has, err := dao.GetUser(&bean.User{UserName: register.UserName})
	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
	}
	if has {
		return nil, ssoerror.ErrorRegisterUserExist
	}

	verify := &bean.Verify{Type: bean.VerifyTypeRegisterEmail, VerifyId: register.Email, Code: register.VerifyCode}
	has, err = dao.GetVerify(verify)
	if err != nil {
		log.Println(err.Error())
		return nil, ssoerror.ErrorInternalServerError
	}

	if !has {
		log.Println("没有查询到数据")
		return nil, ssoerror.ErrorRegisterErrorVerifyCode
	}

	timeNow := time.Now()
	if timeNow.After(*verify.ExpiredTime) {
		log.Println("验证码超时")
		return nil, ssoerror.ErrorRegisterErrorVerifyCode
	}
	if !strings.EqualFold(verify.Code, register.VerifyCode) {
		log.Println("验证码错误")
		return nil, ssoerror.ErrorRegisterErrorVerifyCode
	}

	user := &bean.User{
		UserName:         register.UserName,
		Password:         register.Password,
		Email:            register.Email,
		CreateTime:       &timeNow,
		UpdateTime:       &timeNow,
		IsEmailConfirmed: true,
	}

	_, err = dao.InsertUser(user)

	if err != nil {
		return nil, ssoerror.ErrorInternalServerError
	}

	return user, nil
}
