package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"sso/bean"
	"sso/service"
	"sso/util"
)

func UserRegister(c echo.Context) error {
	user := new(bean.User)
	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	if error := c.Bind(user); error != nil {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}
	if len(user.Name) == 0 || len(user.Password) == 0 {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}

	user, err := service.UserRegister(user)

	if err != nil {
		c.Logger().Debug("error:" + err.Error())
		response.Code = util.StatusResourceExist
		response.Desc = util.StatusText(util.StatusResourceExist)
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
		if err == service.ErrorServiceNotFound {
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

	user := new(bean.User)
	response := bean.Response{Code: util.StatusOK, Desc: util.StatusText(util.StatusOK)}

	if err := c.Bind(user); err != nil {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}
	if len(user.Name) == 0 || len(user.Password) == 0 {
		response.Code = util.StatusIllegalParam
		response.Desc = util.StatusText(util.StatusIllegalParam)
		return c.JSON(http.StatusOK, response)
	}
	token, user, err := service.UserLogin(user)

	if err != nil {
		c.Logger().Debug("error:" + err.Error())
	}

	if err != nil || user == nil {
		response.Code = util.StatusResourceNoExist
		response.Desc = util.StatusText(util.StatusResourceNoExist)
		return c.JSON(http.StatusOK, response)
	}

	response.Data = map[string]string{
		"id":    user.ID.Hex(),
		"token": token.Token,
	}

	return c.JSON(http.StatusOK, response)
}
