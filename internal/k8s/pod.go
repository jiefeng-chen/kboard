package k8s

import (
	"fmt"
	"io/ioutil"
	"kboard/config"
	"kboard/internal"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type IPod interface {
	Create(string, []byte) (err *internal.HttpError)
	Replace(string, string, []byte) (err *internal.HttpError)
	Read(string, string) (*simplejson.Json, *internal.HttpError)
	Delete(string, string) (err *internal.HttpError)
	WriteToEtcd(string, string, []byte) *internal.HttpError
	List(string) (*simplejson.Json, *internal.HttpError)
	Log(string, string) []byte
}

type Pod struct {
	internal.K8sCore
}

func NewPod(Config config.IConfig) IPod {
	return &Pod{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_POD,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/pods/%s",
				Create: "/api/v1/namespaces/%s/pods",
				List:   "/api/v1/namespaces/%s/pods",
				Watch:  "/apis/batch/v1/watch/namespaces/%s/jobs/%s",
			},
		},
	}
}

func (l *Pod) List(ns string) (*simplejson.Json, *internal.HttpError) {
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
