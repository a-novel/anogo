package tests

import (
	"fmt"
	"github.com/a-novel/anogo"
	"reflect"
	"testing"
)

func flatMapEqual(a, b map[string]interface{}) (string, bool) {
	if len(a) != len(b) {
		return fmt.Sprintf("a has %v keys but b has %v", len(a), len(b)), false
	}

	for k, v := range a {
		bv, ok := b[k]
		if !ok {
			return fmt.Sprintf("%s is missing in b", k), false
		}

		if v == nil {
			if bv == nil {
				continue
			}

			return fmt.Sprintf("%s is nil in a but has value %v in b", k, bv), false
		}

		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			if !anogo.SliceEqual(v, bv) {
				return fmt.Sprintf("%s has value %v in a but %v in b", k, v, bv), false
			}
		default:
			if !reflect.DeepEqual(v, bv) {
				return fmt.Sprintf("%s has value %v in a but %v in b", k, v, bv), false
			}
		}
	}

	return "", true
}

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

func TestFlatten(t *testing.T) {
	type subStruct struct {
		Foo string `json:"foo"`
		Qux []string `json:"qux"`
	}

	type mainStruct struct {
		SKey string `json:"s_key"`
		IKey float64 `json:"i_key"`
		MKey map[string]interface{} `json:"m_key"`
		MKeyN map[string]interface{} `json:"m_key_n"`
		StKey subStruct `json:"st_key"`
	}

	strKey := "john"
	subsKey := subStruct{
		Foo: "barbar",
	}

	val := mainStruct{
		SKey: "hello world",
		IKey: 123456,
		MKey: map[string]interface{}{
			"name": &strKey,
			"lastonline": 123456,
			"userSt": &subsKey,
		},
		StKey: subStruct{
			Foo: "bar",
			Qux: []string{"quux", "quuux"},
		},
	}

	out, err := anogo.Flatten(val)
	if err != nil {
		t.Errorf("cannot flatten valid structure : %s", err.Error())
	}

	key, ok := flatMapEqual(out, map[string]interface{}{
		"s_key": "hello world",
		"i_key": float64(123456),
		"m_key_n": nil,
		"st_key.foo": "bar",
		"st_key.qux": []string{"quux", "quuux"},
		"m_key.name": "john",
		"m_key.lastonline": float64(123456),
		"m_key.userSt.foo": "barbar",
		"m_key.userSt.qux": nil,
	})

	if !ok {
		t.Errorf("non matching values : %s", key)
	}
}