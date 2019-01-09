package rest

import (
	"net/http"
	"net/url"
)

type IHttpClient interface {
	Get()
	Post()
	Put()
	Delete()
	Patch()
}

type HttpClient struct {
	baseUrl *url.URL

	header *http.Header

	Client *http.Client
}

func (c *HttpClient) Get() {

}

func (c *HttpClient) Post() {

}

func (c *HttpClient) Delete() {

}

func (c *HttpClient) Put() {

}

func (c *HttpClient) Patch() {

}

func NewHttpClient(url *url.URL, client *http.Client, header *http.Header) IHttpClient {
	return &HttpClient{
		baseUrl: url,
		header:  header,
		Client:  client,
	}
}
