package anogo

import (
	"github.com/a-novel/errors"
	"reflect"
)

/*
	Check if value is a pointer.

	May return:
		- ErrIsNotPtr
		- ErrIsNestedPtr
*/
func IsPtr(ptr interface{}) *errors.Error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return errors.New(ErrIsNotPtr, "ptr is not a valid pointer value, but a direct copy")
	}

	if reflect.ValueOf(ptr).Elem().Kind() == reflect.Ptr {
		return errors.New(ErrIsNestedPtr, "nested pointers are not accepted")
	}

	return nil
}

/*
	Check if slice is a slice of pointers.

	May return:
		- ErrIsNotSlice
		- ErrIsNotPtr
		- ErrIsNestedPtr
*/
func IsSlicePtr(ptr interface{}) *errors.Error {
	if reflect.TypeOf(ptr).Kind() != reflect.Slice {
		return errors.New(ErrIsNotSlice, "non slice value was given to IsSlicePtr")
	}

	if reflect.TypeOf(ptr).Elem().Kind() != reflect.Ptr {
		return errors.New(ErrIsNotPtr, "slice type is not a ptr")
	}

	if reflect.TypeOf(ptr).Elem().Elem().Kind() == reflect.Ptr {
		return errors.New(ErrIsNestedPtr, "nested pointers are not accepted")
	}

	return nil
}
