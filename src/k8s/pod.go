package k8s

import (
	"core"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
	"io/ioutil"
	"net/http"
	"resource"
)

type IPod interface {
	IK8sCore
	List(string) (*simplejson.Json, error)
	Log(string, string) []byte
}

type Pod struct {
	K8sCore
}

func NewPod(Config *config.Context) *Pod {
	return &Pod{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_POD,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/pods/%s",
				Create: "/api/v1/namespaces/%s/pods",
			},
		},
	}
}

func (l *Pod) List(ns string) (*simplejson.Json, *HttpError) {
	url := fmt.Sprintf(l.Urls.Create, ns)
	jsonData := l.get(url)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return jsonData, err
}

func (l *Pod) Log(ns string, name string) []byte {
	url := l.baseApi() + fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/log", ns, name)
	//log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		core.ERROR.Println(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}
