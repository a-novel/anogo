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
