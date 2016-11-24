package controller

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"sso/bean"
	ssoerror "sso/error"
	"sso/request"
	ssoresponse "sso/response"
	"sso/service"
)

func UserRegisetrEMailVerifyCode(c echo.Context) error {

	user := new(bean.User)

	if err := c.Bind(user); err != nil {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if err := service.UserRegisetrEMailVerifyCode(user); err != nil {
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func UserInfo(c echo.Context) error {

	token := c.Request().Header().Get("token")
	userBean, err := service.UserInfoByToken(token)
	if err != nil {
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	responseData := &ssoresponse.User{
		UserName:         userBean.UserName,
		Name:             userBean.Name,
		Icon:             userBean.Icon,
		Sex:              userBean.Sex,
		NickName:         userBean.NickName,
		Email:            userBean.Email,
		Mobile:           userBean.Mobile,
		IsEmailConfirmed: userBean.IsEmailConfirmed,
	}
	if userBean.Birthday != nil {
		birthday := userBean.Birthday.Unix()
		responseData.Birthday = &birthday
	}

	response.Data = responseData
	return c.JSON(http.StatusOK, response)
}

func UserUpdate(c echo.Context) error {

	token := c.Request().Header().Get("token")

	update := &request.Update{}

	if err := c.Bind(update); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if err := service.UserUpdate(token, update); err != nil {
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func ChangeEmailVerifyCode(c echo.Context) error {

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func ChangeEmail(c echo.Context) error {

	type ChangeEmail struct {
		NewMail    string `json:"newEmail,omitempty" form:"newEmail"`
		VerifyCode string `json:"verifyCode,omitempty" form:"verifyCode"`
	}
	token := c.Request().Header().Get("token")
	changeEmail := &ChangeEmail{}

	if err := c.Bind(changeEmail); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if err := service.ChangeEmail(token, changeEmail.NewMail, changeEmail.VerifyCode); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func UserChangePassword(c echo.Context) error {

	type ChangePassword struct {
		OriginalPassword string `json:"originalpassword,omitempty" form:"originalpassword"`
		NewPassword      string `json:"newpassword,omitempty" form:"newpassword"`
	}
	changePassoword := &ChangePassword{}

	token := c.Request().Header().Get("token")

	if err := c.Bind(changePassoword); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if err := service.UserChangePassword(token, changePassoword.OriginalPassword, changePassoword.NewPassword); err != nil {
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func UserRegister(c echo.Context) error {

	register := &request.Register{}
	if err := c.Bind(register); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	_, err := service.UserRegister(register)

	if err != nil {
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func UserLogout(c echo.Context) error {

	token := c.Request().Header().Get("token")

	if len(token) == 0 {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	err := service.UserLogout(token)
	if err != nil {
		return ControllerHandleError(c, err)
	}
	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func UserLogin(c echo.Context) error {

	login := &request.Login{}
	//用户信息
	if err := c.Bind(login); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	userBean, tokenBean, err := service.UserLogin(login)
	if err != nil {
		c.Logger().Debug("error:" + err.Error())
		return ControllerHandleError(c, err)
	}
	response := bean.NEWResponse(ssoerror.CommonSuccess)
	response.Data = map[string]interface{}{
		"id":    userBean.ID,
		"token": tokenBean.Token,
	}

	return c.JSON(http.StatusOK, response)
}
