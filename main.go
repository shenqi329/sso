package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"sso/controller"
	"sso/mysql"
)

func defaultServer(c echo.Context) error {
	return c.JSON(http.StatusOK, "a message from sso")
}

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	e := echo.New()

	e.Use(echomiddleware.Logger())

	e.Post("/user/info", controller.UserInfo)
	e.POST("/user/register", controller.UserRegister)
	e.POST("/user/login", controller.UserLogin)
	e.POST("/user/logout", controller.UserLogout)

	e.GET("/", defaultServer)

	mysql.Connect()
	mysql.GetXormEngine()
	e.SetDebug(true)
	log.Println("sso run on port:8081")
	e.Run(standard.New(":8081"))
}
