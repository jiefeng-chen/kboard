package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"kboard/internal"
)

type IIngress interface {
	Create(string, []byte) (err *internal.HttpError)
	Replace(string, string, []byte) (err *internal.HttpError)
	Read(string, string) (*simplejson.Json, *internal.HttpError)
	Delete(string, string) (err *internal.HttpError)
	WriteToEtcd(string, string, []byte) *internal.HttpError
	List(ns string) (*simplejson.Json, *internal.HttpError)
}

type Ingress struct {
	internal.K8sCore
}

func NewIngress(Config config.IConfig) IIngress {
	return &Ingress{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_INGRESS,
			Urls: internal.Urls{
				Read:   "/apis/extensions/v1beta1/namespaces/%s/ingresses/%s",
				Create: "/apis/extensions/v1beta1/namespaces/%s/ingresses",
			},
		},
	}
}

func (l *Ingress) List(ns string) (*simplejson.Json, *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.Create, ns)
	jsonData := l.get(url)
	httpResult := internal.GetHttpCode(jsonData)
	err := internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind+"List" {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return jsonData, err
}
