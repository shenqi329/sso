package controller

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"sso/bean"
	ssoerror "sso/error"
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

	userBean.Password = ""
	response := bean.NEWResponse(ssoerror.CommonSuccess)
	response.Data = map[string]interface{}{
		"userinfo": userBean,
	}
	return c.JSON(http.StatusOK, response)
}

func UserUpdate(c echo.Context) error {

	token := c.Request().Header().Get("token")
	response := bean.NEWResponse(ssoerror.CommonSuccess)

	user := new(bean.User)
	if err := c.Bind(user); err != nil {
		log.Println(err.Error())
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if err := service.UserUpdate(token, user); err != nil {
		return ControllerHandleError(c, err)
	}

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

	register := &bean.Register{}
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

	login := &bean.Login{}
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
