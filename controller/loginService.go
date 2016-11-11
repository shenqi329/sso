package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"sso/bean"
)

func Login(c echo.Context) error {

	user := new(bean.UserBean)
	if error := c.Bind(user); error != nil {
		return error
	}
	fmt.Println("username:" + user.Name + "password" + user.Password)

	// if !service.UserRegister(user) {
	// 	responseBean := bean.ResponseBean{Code: "000002", Desc: "resource is exist"}
	// 	return c.JSON(http.StatusOK, responseBean)
	// }

	responseBean := bean.ResponseBean{Code: "000001", Desc: "success"}
	return c.JSON(http.StatusOK, responseBean)
}
