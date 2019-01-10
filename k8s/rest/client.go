package rest

import (
	"net/http"
	"net/url"
)

type IRESTClient interface {
	Get()
	Post()
	Put()
	Delete()
	Patch()
	Do()
}

type RESTClient struct {
	baseUrl *url.URL

	header *http.Header

	Client *http.Client
}

func (c *RESTClient) getUrl() string {

	return ""
}

func (c *RESTClient) Get() {

}

func (c *RESTClient) Post() {

}

func (c *RESTClient) Delete() {

}

func (c *RESTClient) Put() {

}

func (c *RESTClient) Patch() {

}

func (c *RESTClient) Do() {

}

func NewHttpClient(url *url.URL, header *http.Header) IRESTClient {
	return &RESTClient{
		baseUrl: url,
		header:  header,
		Client:  &http.Client{},
	}
}
