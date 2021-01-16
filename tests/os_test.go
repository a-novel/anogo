package tests

import (
	"github.com/a-novel/anogo"
	"strings"
	"testing"
)

func TestGetExecPath(t *testing.T) {
	expath, err := anogo.GetExecPath()
	if err != nil {
		t.Fatalf("cannot get executable file path : %s", err.Error())
	}

	if !strings.HasSuffix(expath, "/tests") {
		t.Errorf("unexpected executable path : should end with /tests, got %s", expath)
	}
}
