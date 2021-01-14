package tests

import (
	"github.com/a-novel/anogo"
	"reflect"
	"testing"
)

func TestToMapInterface(t *testing.T) {
	val1 := map[string]interface{}{
		"foo":       "bar",
		"timestamp": float64(123456),
	}

	val2 := struct {
		Foo       string `json:"foo"`
		Timestamp int    `json:"timestamp"`
	}{Foo: "bar", Timestamp: 123456}

	val3 := []string{"hello", "world"}

	mval1, err := anogo.ToMapInterface(val1)
	if err != nil {
		t.Errorf("cannot convert map[string]interface{} to itself : %s", err.Error())
	} else if !reflect.DeepEqual(mval1, val1) {
		t.Errorf("unexpected value returned from val1 conversion : got %v instead of %v", mval1, val1)
	}

	mval2, err := anogo.ToMapInterface(val2)
	if err != nil {
		t.Errorf("cannot convert struct to map[string]interface{} : %s", err.Error())
	} else if !reflect.DeepEqual(mval2, val1) {
		t.Errorf("unexpected value returned from val2 conversion : got %v instead of %v", mval2, val1)
	}

	if _, err := anogo.ToMapInterface(val3); err == nil {
		t.Errorf("no error returned when converting non map type to map[string]interface{}")
	} else if err.ID != anogo.ErrIsNotMappable {
		t.Errorf("unexpected error ID when converting non map type : got %s, expected %s", err.ID, anogo.ErrIsNotMappable)
	}
}
