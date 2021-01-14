package anogo

import (
	"github.com/a-novel/errors"
	"reflect"
)

func IsPtr(ptr interface{}) *errors.Error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return errors.New(ErrIsNotPtr, "ptr is not a valid pointer value, but a direct copy")
	}

	if reflect.Indirect(reflect.ValueOf(ptr)).Kind() == reflect.Ptr {
		return errors.New(ErrIsNestedPtr, "nested pointers are not accepted")
	}

	return nil
}
