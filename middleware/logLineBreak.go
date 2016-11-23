package middleware

import (
	"fmt"
	"github.com/labstack/echo"
)

func LogLineBreak() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			fmt.Println("\r\n\r\n")
			if err = next(c); err != nil {
				c.Error(err)
			}

			return err
		}
	}
}
