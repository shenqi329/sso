package controller

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"regexp"
	"sso/bean"
	ssoerror "sso/error"
	"sso/service"
)

func UserRegisetrEMailVerifyCode(c echo.Context) error {

	user := new(bean.User)

	if err := c.Bind(user); err != nil {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}
	log.Println(bean.StructToJsonString(user))

	if len(user.UserName) == 0 || len(user.Email) == 0 {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	regexp.MustCompile(pattern)
	match, err := regexp.MatchString(pattern, user.Email)
	if !match || err != nil {
		return ControllerHandleError(c, ssoerror.ErrorRegisterEmailFormat)
	}

	err = service.UserRegisetrEMailVerifyCode(user)
	if err != nil {
		return ControllerHandleError(c, err)
	}

	response := bean.NEWResponse(ssoerror.CommonSuccess)
	return c.JSON(http.StatusOK, response)
}

func UserInfo(c echo.Context) error {

	token := c.Request().Header().Get("token")

	userBean, err := service.UserInfo(token)
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

func UserRegister(c echo.Context) error {
	user := new(bean.User)
	email := new(bean.Email)

	if err := c.Bind(user); err != nil {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}
	if len(user.UserName) == 0 ||
		len(user.Password) == 0 ||
		len(user.Email) == 0 {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if err := c.Bind(email); err != nil {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}
	if len(email.UserName) == 0 ||
		len(email.Email) == 0 ||
		len(email.Code) == 0 {
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	log.Println(bean.StructToJsonString(email))

	user, err := service.UserRegister(user, email)
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

	userBean := &bean.User{}
	tokenBean := &bean.Token{}

	//检测用户信息
	if err := c.Bind(userBean); err != nil {
		log.Println("获取用户信息失败")
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}
	if len(userBean.UserName) <= 0 ||
		len(userBean.Password) <= 0 {
		log.Println("获取用户信息失败")
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	//检查授权信息
	if err := c.Bind(tokenBean); err != nil {
		log.Println("获取授权信息失败")
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	if len(tokenBean.DeviceId) <= 0 ||
		len(tokenBean.AppId) <= 0 ||
		len(tokenBean.Platform) <= 0 {
		log.Println("获取授权信息失败")
		return ControllerHandleError(c, ssoerror.ErrorIllegalParams)
	}

	userBean, tokenBean, err := service.UserLogin(userBean, tokenBean)
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
