package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	"net/http"
	"sso/controller"
)

func defaultServer(c echo.Context) error {
	return c.JSON(http.StatusOK, "a message from sso")
}

func main() {
	e := echo.New()

	e.Use(echomiddleware.Logger())

	e.POST("/user/login", controller.UserLogin)
	e.POST("/user/register", controller.UserRegister)
	e.POST("/user/logout", controller.UserLogout)

	e.GET("/", defaultServer)

	e.SetDebug(true)

	fmt.Println("server run on port:80")
	e.Run(standard.New(":80"))
}
