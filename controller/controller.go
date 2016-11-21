package controller

import (
	"log"
	"reflect"
)

func ControllerHandleError(err error) error {
	log.Println(err.Error())

	typ := reflect.TypeOf(err)

	if typ == reflect.Ptr {
		typ = typ.Elem()
	}

	return err
}
