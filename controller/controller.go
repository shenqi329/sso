package controller

import (
	"github.com/labstack/echo"
	//"log"
	"net/http"
	"reflect"
	"sso/bean"
	ssoerror "sso/error"
)

func ControllerHandleError(c echo.Context, err error) error {

	e := reflect.Indirect(reflect.ValueOf(err))

	response := &bean.Response{}

	code := e.FieldByName("Code")
	if code.Kind() == reflect.String {
		response.Code = code.String()
	}
	desc := e.FieldByName("Desc")
	if desc.Kind() == reflect.String {
		response.Desc = desc.String()
	}
	//log.Println(bean.StructToJsonString(response))

	if len(response.Code) <= 0 {
		response = bean.NEWResponse(ssoerror.CommonInternalServerError)
	}
	return c.JSON(http.StatusOK, response)
}
