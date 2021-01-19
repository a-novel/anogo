package anogo

import (
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"reflect"
)

type ToMapOptions struct {
	PreservePointers bool
	PreserveStructs bool
}

/*
	Convert v to the underlying value of ptr.

	May return:
		- ErrIsNotPtr
		- ErrIsNestedPtr
		- ErrIsNotMappable
		- ErrCannotMarshal
		- ErrCannotUnmarshal
*/
func ToMap(v interface{}, ptr interface{}) *errors.Error {
	if err := IsPtr(ptr); err != nil {
		return err
	}

	kind := reflect.TypeOf(v).Kind()

	if kind != reflect.Map && kind != reflect.Struct {
		return errors.New(
			ErrIsNotMappable,
			fmt.Sprintf("go type %s cannot be parsed as map", kind.String()),
		)
	}

	jsonString, err := json.Marshal(v)
	if err != nil {
		return errors.New(
			ErrCannotMarshal,
			fmt.Sprintf("cannot marshal value : %s", err.Error()),
		)
	}

	if err := json.Unmarshal(jsonString, ptr); err != nil {
		return errors.New(
			ErrCannotUnmarshal,
			fmt.Sprintf("cannot unmarshal value : %s", err.Error()),
		)
	}

	return nil
}

/*
	Convert v to a map[string]interface{}.

	May return:
		- ErrIsNotPtr
		- ErrIsNestedPtr
		- ErrIsNotMappable
		- ErrCannotMarshal
		- ErrCannotUnmarshal
*/
func ToMapInterface(v interface{}) (map[string]interface{}, *errors.Error) {
	if mv, ok := v.(map[string]interface{}); ok {
		return mv, nil
	}

	var output map[string]interface{}
	if err := ToMap(v, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Convert v to a map[string]string.

	May return:
		- ErrIsNotPtr
		- ErrIsNestedPtr
		- ErrIsNotMappable
		- ErrCannotMarshal
		- ErrCannotUnmarshal
*/
func ToMapString(v interface{}) (map[string]string, *errors.Error) {
	if mv, ok := v.(map[string]string); ok {
		return mv, nil
	}

	var output map[string]string
	if err := ToMap(v, &output); err != nil {
		return nil, err
	}

	return output, nil
}

/*
	Flattens a map-like object.

	May return:
		- ErrIsNotPtr
		- ErrIsNestedPtr
		- ErrIsNotMappable
		- ErrCannotMarshal
		- ErrCannotUnmarshal
*/
func Flatten(data interface{}) (map[string]interface{}, *errors.Error) {
	if data == nil {
		return nil, nil
	}

	output := map[string]interface{}{}

	mdata, err := ToMapInterface(data)
	if err != nil {
		return nil, err
	}

	for k, v := range mdata {
		if v == nil {
			output[k] = v
			continue
		}

		for reflect.TypeOf(v).Kind() == reflect.Ptr {
			v = reflect.ValueOf(v).Elem().Interface()
		}

		kind := reflect.TypeOf(v).Kind()

		if kind == reflect.Map || kind == reflect.Struct {
			flattened, err := Flatten(v)
			if err != nil {
				return nil, err
			}

			for sk, sv := range flattened {
				output[fmt.Sprintf("%s.%s", k, sk)] = sv
			}
		} else {
			output[k] = v
		}
	}

	return output, nil
}
