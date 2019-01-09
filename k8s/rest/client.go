package rest

type IHttpClient interface {
	Get()
	Post()
	Put()
	Delete()
	Replace()
}
