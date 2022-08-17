package httpClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const timeout = 30

type Client interface {
	GET(get Get) (response []byte, err error)
	POST(post Post) (response []byte, err error)
}

func NewHttpClient() Client {
	return &client{
		httpClient: http.Client{Timeout: time.Second * timeout},
	}
}

type client struct {
	httpClient http.Client
}

func (r *client) GET(get Get) (response []byte, err error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
	)

	rawUrl := get.host + get.path + get.queryString

	if httpRequest, err = http.NewRequest("GET", rawUrl, nil); err != nil {
		return nil, err
	}

	defer closeHttpBody(httpResponse)
	if httpResponse, err = r.httpClient.Do(httpRequest); err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s:%s", "Status Code", httpResponse.StatusCode))
	}

	if response, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *client) POST(post Post) (response []byte, err error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		requestBytes []byte
	)

	if requestBytes, err = json.Marshal(post.request); err != nil {
		return nil, err
	}

	rawUrl := post.host + post.path
	if httpRequest, err = http.NewRequest("POST", rawUrl, bytes.NewBuffer(requestBytes)); err != nil {
		return nil, err
	}

	defer closeHttpBody(httpResponse)
	if httpResponse, err = r.httpClient.Do(httpRequest); err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s:%s", "Status Code", httpResponse.StatusCode))
	}

	if response, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}

	return response, nil
}

func closeHttpBody(httpResponse *http.Response) {
	if httpResponse != nil && httpResponse.Body != nil {
		httpResponse.Body.Close()
	}
}
