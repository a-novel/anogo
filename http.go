package anogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"io/ioutil"
	"net/http"
)

func PostWithHeader(url string, payload interface{}, headers map[string]string) (int, []byte, *errors.Error) {
	mrsh, err := json.Marshal(payload)
	if err != nil {
		return 0, nil, errors.New(ErrCannotMarshal, fmt.Sprintf("cannot marshal payload : %s", err.Error()))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mrsh))
	if err != nil {
		return 0, nil, errors.New(ErrCannotCreatePostRequest, fmt.Sprintf("cannot create post request : %s", err.Error()))
	}

	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, errors.New(ErrRequestFailed, fmt.Sprintf("request failed : %s", err.Error()))
	}

	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, errors.New(ErrCannotParseRequestBody, fmt.Sprintf("cannot parse requets body : %s", err.Error()))
	}

	return resp.StatusCode, output, nil
}

func PostJSON(url string, payload interface{}) (int, []byte, *errors.Error) {
	return PostWithHeader(url, payload, map[string]string{"Content-Type": "application/json"})
}
