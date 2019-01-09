package rest

import (
	"net/url"
	"net/http"
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


func NewHttpClient(url *url.URL, client *http.Client) IHttpClient {
	return &HttpClient{
		baseUrl: url,
		Client: client,
	}
}