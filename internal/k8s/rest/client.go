package rest

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

type IHttpClient interface {
	Get() (resp *http.Response, err error)
	Post() (resp *http.Response, err error)
	Put() (resp *http.Response, err error)
	Delete() (resp *http.Response, err error)
	Patch() (resp *http.Response, err error)
	Do(method string) (resp *http.Response, err error)
	SetHeader(key string, values ...string)

	// 非restful实现
	Watch() (resp *http.Response, err error)
	List() (resp *http.Response, err error)
}

type HttpClient struct {
	baseUrl *url.URL

	headers http.Header

	namespace    string
	resourceType string
	resourceName string

	timeout time.Duration

	err  error
	body io.Reader

	ctx context.Context

	Client *http.Client
}

func (c *HttpClient) getUrl() string {
	// 拼接地址

	return ""
}

func (c *HttpClient) Get() (resp *http.Response, err error) {
	return c.Do("GET")
}

func (c *HttpClient) Post() (resp *http.Response, err error) {
	return c.Do("POST")
}

func (c *HttpClient) Delete() (resp *http.Response, err error) {
	return c.Do("DELETE")
}

func (c *HttpClient) Put() (resp *http.Response, err error) {
	return c.Do("PUT")
}

func (c *HttpClient) Patch() (resp *http.Response, err error) {
	return c.Do("PATCH")
}

func (c *HttpClient) Do(method string) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, c.getUrl(), c.body)
	if err != nil {
		return nil, err
	}
	// 设置header头部
	req.Header = c.headers
	// 用json数据格式
	req.Header.Add("Content-Type", "application/json")

	return c.Client.Do(req)
}

func (c *HttpClient) SetHeader(key string, values ...string) {
	if c.headers == nil {
		c.headers = http.Header{}
	}
	c.headers.Del(key)
	for _, header := range values {
		c.headers.Add(key, header)
	}
}

func (c *HttpClient) Watch() (resp *http.Response, err error) {

}

func (c *HttpClient) List() (resp *http.Response, err error) {

}

func NewHttpClient(url *url.URL, headers http.Header) IHttpClient {
	return &HttpClient{
		baseUrl: url,
		headers: headers,
		Client:  &http.Client{},
	}
}
