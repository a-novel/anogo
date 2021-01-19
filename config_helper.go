package anogo

import (
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"io/ioutil"
)

/*
	Reads a config file and assign it to a pointer.

	May return:
		- ErrCannotOpenConfigFile
		- ErrCannotParseConfigFile
*/
func ReadConfig(cpath string, ptr interface{}) *errors.Error {
	file, err := ioutil.ReadFile(cpath)
	if err != nil {
		return errors.New(
			ErrCannotOpenConfigFile,
			fmt.Sprintf("cannot read config file : %s", err.Error()),
		)
	}

	if err = json.Unmarshal(file, &ptr); err != nil {
		return errors.New(
			ErrCannotParseConfigFile,
			fmt.Sprintf("cannot parse config file : %s", err.Error()),
		)
	}

	return nil
}
