package middleware

import (
	"github.com/labstack/echo"
	"strings"
)

func LoginFilter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			strings.EqualFold(s, t)
			req.URI()
		}
	}
}
