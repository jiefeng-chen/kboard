package k8s

import (
	"fmt"
	"io/ioutil"
	"kboard/config"
	"kboard/k8s/resource"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type IPod interface {
	Create(string, []byte) (err *HttpError)
	Replace(string, string, []byte) (err *HttpError)
	Read(string, string) (*simplejson.Json, *HttpError)
	Delete(string, string) (err *HttpError)
	WriteToEtcd(string, string, []byte) *HttpError
	List(string) (*simplejson.Json, *HttpError)
	Log(string, string) []byte
}

type Pod struct {
	K8sCore
}

func NewPod(Config config.IConfig) IPod {
	return &Pod{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_POD,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/pods/%s",
				Create: "/api/v1/namespaces/%s/pods",
				List: "/api/v1/namespaces/%s/pods",
				Watch: "/apis/batch/v1/watch/namespaces/%s/jobs/%s",
			},
		},
	}
}

func (l *Pod) List(ns string) (*simplejson.Json, *HttpError) {
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

func (l *Pod) Log(ns string, name string) []byte {
	url := l.baseApi() + fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/log?pretty=true&tailLines=500", ns, name)
	//log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}
