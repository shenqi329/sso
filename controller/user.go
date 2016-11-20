package controller

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"sso/bean"
	"sso/service"
	"sso/util"
)

func UserRegisetrEMailVerifyCode(c echo.Context) error {

	user := new(bean.User)
	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	if err := c.Bind(user); err != nil {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		log.Println(err.Error())
		return c.JSON(http.StatusOK, response)
	}

	log.Println(bean.StructToJsonString(user))

	if len(user.UserName) == 0 || len(user.Email) == 0 {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	err := service.UserRegisetrEMailVerifyCode(user)

	if err != nil {
		if err == service.ErrorDatabaseOperation {
			response.Code = util.StatusIllegalParam
			response.Desc = util.StatusText(util.StatusInnerError)

		} else if err == service.ErrorResourceExist {
			response.Code = util.StatusResourceExist
			response.Desc = util.StatusText(util.StatusResourceExist)
		}
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}

func UserInfo(c echo.Context) error {

	token := c.Request().Header().Get("token")

	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	userBean, err := service.UserInfo(token)

	if err != nil {
		if err == service.ErrorNotFound {
			response.Code = util.StatusResourceNoExist
			response.Desc = util.StatusText(util.StatusResourceNoExist)
		} else if err == service.ErrorParams {
			response.Code = util.StatusIllegalParam
			response.Desc = util.StatusText(util.StatusIllegalParam)
		} else {
			response.Code = util.StatusInnerError
			response.Desc = util.StatusText(util.StatusInnerError)
		}
	} else {
		userBean.Password = ""
		response.Data = map[string]interface{}{
			"userinfo": userBean,
		}
	}

	return c.JSON(http.StatusOK, response)
}

func UserRegister(c echo.Context) error {
	user := new(bean.User)
	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	if error := c.Bind(user); error != nil {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}
	if len(user.UserName) == 0 || len(user.Password) == 0 {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	user, err := service.UserRegister(user)

	if err != nil {
		if err == service.ErrorResourceExist {
			response.Code = util.StatusResourceExist
			response.Desc = util.StatusText(util.StatusResourceExist)
		} else {
			response.Code = util.StatusInnerError
			response.Desc = util.StatusText(util.StatusInnerError)
		}
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}

func UserLogout(c echo.Context) error {

	token := c.Request().Header().Get("token")
	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	if len(token) == 0 {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	err := service.UserLogout(token)
	if err != nil {
		if err == service.ErrorNotFound {
			response.Code = util.StatusResourceNoExist
			response.Desc = util.StatusText(util.StatusResourceNoExist)
		} else {
			response.Code = util.StatusInnerError
			response.Desc = util.StatusText(util.StatusInnerError)
		}
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, response)
}

func UserLogin(c echo.Context) error {

	userBean := &bean.User{}
	tokenBean := &bean.Token{}

	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	//检测用户信息
	if err := c.Bind(userBean); err != nil {
		log.Println("获取用户信息失败")
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}
	if len(userBean.UserName) <= 0 ||
		len(userBean.Password) <= 0 {
		log.Println("获取用户信息失败")
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	//检查授权信息
	if err := c.Bind(tokenBean); err != nil {
		log.Println("获取授权信息失败")
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	if len(tokenBean.DeviceId) <= 0 ||
		len(tokenBean.AppId) <= 0 ||
		len(tokenBean.Platform) <= 0 {
		log.Println("获取授权信息失败")
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	userBean, tokenBean, err := service.UserLogin(userBean, tokenBean)
	if err != nil {
		c.Logger().Debug("error:" + err.Error())
		if err == service.ErrorNotFound {
			response.Code = util.StatusResourceNoExist
			response.Desc = util.StatusText(util.StatusResourceNoExist)
		} else {
			response.Code = util.StatusInnerError
			response.Desc = util.StatusText(util.StatusInnerError)
		}
		return c.JSON(http.StatusOK, response)
	}

	response.Data = map[string]interface{}{
		"id":    userBean.ID,
		"token": tokenBean.Token,
	}

	return c.JSON(http.StatusOK, response)
}
