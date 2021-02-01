package anogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"io/ioutil"
	"net/http"
)

type Header map[string]string

func (h Header) JSON() {
	h["Content-Type"] = "application/json"
}

func (h Header) Token(token string) {
	h["Authorization"] = token
}

func (h Header) UserAgent(ua string) {
	h["User-Agent"] = ua
}

func Request(method, url string, payload interface{}, headers Header) (int, []byte, *errors.Error) {
	var mrsh []byte
	var err error

	if payload != nil {
		mrsh, err = json.Marshal(payload)
		if err != nil {
			return 0, nil, errors.New(ErrCannotMarshal, fmt.Sprintf("cannot marshal payload : %s", err.Error()))
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(mrsh))
	if err != nil {
		return 0, nil, errors.New(ErrCannotCreateRequest, fmt.Sprintf("cannot create request : %s", err.Error()))
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

func Post(url string, payload interface{}, headers Header) (int, []byte, *errors.Error) {
	return Request(http.MethodPost, url, payload, headers)
}

func Get(url string, headers Header) (int, []byte, *errors.Error) {
	return Request(http.MethodGet, url, nil, headers)
}

func Delete(url string, payload interface{}, headers Header) (int, []byte, *errors.Error) {
	return Request(http.MethodDelete, url, payload, headers)
}

func Put(url string, payload interface{}, headers Header) (int, []byte, *errors.Error) {
	return Request(http.MethodPut, url, payload, headers)
}