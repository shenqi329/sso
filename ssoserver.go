package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"sso/controller"
	ssomiddleware "sso/middleware"
	"sso/mysql"
)

func defaultServer(c echo.Context) error {
	return c.JSON(http.StatusOK, "a message from sso")
}

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	e := echo.New()

	e.Pre(ssomiddleware.LogLineBreak())
	e.Pre(echomiddleware.Logger())

	e.GET("/user/info", controller.UserInfo)
	e.Post("/user/update", controller.UserUpdate)
	e.Post("/user/update/email", controller.ChangeEmail)
	e.Post("/user/update/email/verifycode", controller.ChangeEmailVerifyCode)
	e.Post("/user/register/changepassword", controller.UserChangePassword)
	e.POST("/user/register/email/verifycode", controller.UserRegisetrEMailVerifyCode)
	e.POST("/user/register", controller.UserRegister)
	e.POST("/user/login", controller.UserLogin)
	e.POST("/user/logout", controller.UserLogout)

	e.GET("/", defaultServer)

	mysql.GetXormEngine()
	e.SetDebug(true)
	log.Println("sso run on port:8081")
	e.Run(standard.New(":8081"))
}
