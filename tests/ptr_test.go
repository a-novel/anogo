package tests

import (
	"github.com/a-novel/anogo"
	"testing"
)

func TestIsPtr(t *testing.T) {
	val := &map[string]interface{}{}

	if err := anogo.IsPtr(*val); err == nil {
		t.Errorf("IsPtr did not returned any error for nil interface")
	} else if err.ID != anogo.ErrIsNotPtr {
		t.Errorf("unexpected error returned with nil interface : got %s, expected %s", err.ID, anogo.ErrIsNotPtr)
	}

	if err := anogo.IsPtr(&val); err == nil {
		t.Errorf("IsPtr did not returned any error for nested pointer")
	} else if err.ID != anogo.ErrIsNestedPtr {
		t.Errorf("unexpected error returned with nil interface : got %s, expected %s", err.ID, anogo.ErrIsNestedPtr)
	}

	if err := anogo.IsPtr(val); err != nil {
		t.Errorf("IsPtr returned unexpected error with valid value : %s", err.Error())
	}
}

func TestIsSlicePtr(t *testing.T) {
	var val []*map[string]interface{}
	var val1 []map[string]interface{}
	var val2 []**map[string]interface{}

	if err := anogo.IsSlicePtr(val1); err == nil {
		t.Errorf("IsSlicePtr did not returned any error for nil interface")
	} else if err.ID != anogo.ErrIsNotPtr {
		t.Errorf("unexpected error returned with nil interface : got %s, expected %s", err.ID, anogo.ErrIsNotPtr)
	}

	if err := anogo.IsSlicePtr(val2); err == nil {
		t.Errorf("IsSlicePtr did not returned any error for nested pointer")
	} else if err.ID != anogo.ErrIsNestedPtr {
		t.Errorf("unexpected error returned with nil interface : got %s, expected %s", err.ID, anogo.ErrIsNestedPtr)
	}

	if err := anogo.IsSlicePtr(val); err != nil {
		t.Errorf("IsSlicePtr returned unexpected error with valid value : %s", err.Error())
	}
}
