package middleware

// import (
// 	"github.com/labstack/echo"
// 	"log"
// )

// func LoginFilter() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) (err error) {
// 			req := c.Request()
// 			//resp := c.Response()

// 			if err = next(c); err != nil {
// 				c.Error(err)
// 			}
// 			log.Println("LoginFilter")
// 			return err
// 		}
// 	}
// }
