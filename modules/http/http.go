package http

import (
	"bytes"
	"encoding/json"
	"github.com/dop251/goja"
	"io"
	"net/http"
)

type Module struct {
	runtime *goja.Runtime
}

type Response struct {
	Body       string
	Headers    map[string][]string
	StatusCode int
}

func (h Module) Get(url string) Response {
	return h.request("GET", url, nil)
}

func (h Module) Head(url string) Response {
	return h.request("HEAD", url, nil)
}

func (h Module) Post(url, contentType string, body interface{}) Response {
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	return h.wrapResponse(resp)
}

func (h Module) Put(url string, body interface{}) Response {
	return h.Request("PUT", url, body)
}

func (h Module) Delete(url string, body interface{}) Response {
	return h.Request("DELETE", url, body)
}

func (h Module) Request(method, url string, reqBody interface{}) Response {
	data, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	return h.request(method, url, bytes.NewBuffer(data))
}

func (h Module) request(method, url string, reqBody io.Reader) Response {
	request, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	return h.wrapResponse(resp)
}

func (h Module) wrapResponse(resp *http.Response) Response {
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return Response{
		Body:       string(body),
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
	}
}