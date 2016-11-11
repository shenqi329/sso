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

	e.GET("/user/login", controller.Login)
	e.POST("/user/login", controller.Login)

	e.GET("/", defaultServer)

	fmt.Println("server run on port:81")
	e.Run(standard.New(":81"))
}
