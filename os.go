package anogo

import (
	"github.com/a-novel/errors"
	"path/filepath"
	"runtime"
)

func GetExecPath() (string, *errors.Error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New(ErrCannotReadFilePath, "cannot read current file path")
	}

	return filepath.Dir(filename), nil
}
