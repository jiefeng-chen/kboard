package k8s

import (
	"fmt"
	"kboard/config"
	"kboard/k8s/resource"

	"github.com/bitly/go-simplejson"
)

type IIngress interface {
	Create(string, []byte) (err *HttpError)
	Replace(string, string, []byte) (err *HttpError)
	Read(string, string) (*simplejson.Json, *HttpError)
	Delete(string, string) (err *HttpError)
	WriteToEtcd(string, string, []byte) *HttpError
	List(ns string) (*simplejson.Json, *HttpError)
}

type Ingress struct {
	K8sCore
}

func NewIngress(Config *config.Config) IIngress {
	return &Ingress{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_INGRESS,
			Urls: Urls{
				Read:   "/apis/extensions/v1beta1/namespaces/%s/ingresses/%s",
				Create: "/apis/extensions/v1beta1/namespaces/%s/ingresses",
			},
		},
	}
}

func (l *Ingress) List(ns string) (*simplejson.Json, *HttpError) {
	url := fmt.Sprintf(l.Urls.Create, ns)
	jsonData := l.get(url)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind+"List" {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return jsonData, err
}
