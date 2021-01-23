package anogo

import (
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"reflect"
)

func StartGinTestServer(routers []func (*gin.Engine)) (*httptest.Server, func(p string) string) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	for _, router := range routers {
		router(r)
	}

	ts := httptest.NewServer(r)
	tsUrl := func(p string) string {
		return fmt.Sprintf("%s%s", ts.URL, p)
	}

	return ts, tsUrl
}

func TestPostJSON(url string, body interface{}, expectedStatus int, expectedResponse map[string]interface{}) *errors.Error {
	st, bb, err := PostJSON(url, body)

	if err != nil {
		return err
	}

	if st != expectedStatus {
		return errors.New(
			ErrUnexpectedResponseStatus,
			fmt.Sprintf("unexpected response status : got %v instead of %v", st, expectedStatus),
		)
	}

	resp := map[string]interface{}{}
	if err := json.Unmarshal(bb, &resp); err != nil {
		return errors.New(ErrCannotUnmarshal, fmt.Sprintf("cannot unmarshal post response : %s\n%s", err.Error(), string(bb)))
	}

	if !reflect.DeepEqual(resp, expectedResponse) {
		return errors.New(
			ErrUnexpectedResponse,
			fmt.Sprintf("unexpected response from request : got %v instead of %v", resp, expectedResponse),
		)
	}

	return nil
}
