package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"kboard/internal"
)

type IReplicationController interface {
	Create(string, []byte) (err *internal.HttpError)
	Replace(string, string, []byte) (err *internal.HttpError)
	Read(string, string) (*simplejson.Json, *internal.HttpError)
	Delete(string, string) (err *internal.HttpError)
	WriteToEtcd(string, string, []byte) *internal.HttpError

	List(string) (*simplejson.Json, *internal.HttpError)
	Scale(string, string, int) *internal.HttpError
	Patch(string, string, []byte) *simplejson.Json
}

type ReplicationController struct {
	internal.K8sCore
}

func NewReplicationController(Config config.IConfig) IReplicationController {
	return &ReplicationController{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_REPLICATION_CONTROLLER,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/replicationcontrollers/%s",
				Create: "/api/v1/namespaces/%s/replicationcontrollers",
			},
		},
	}
}

func (l *ReplicationController) List(ns string) (jsonData *simplejson.Json, err *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns)
	jsonData = l.get(url)
	httpResult := internal.GetHttpCode(jsonData)
	err = internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
		return
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
		return
	}
	return
}

func (l *ReplicationController) Scale(ns string, rc string, num int) *internal.HttpError {
	spec := fmt.Sprintf("{\"spec\":{\"replicas\":%d}}", num)
	json := l.Patch(ns, rc, []byte(spec))
	httpResult := internal.GetHttpCode(json)
	err := internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	return err
}

func (l *ReplicationController) Patch(ns string, rc string, data []byte) *simplejson.Json {
	url := fmt.Sprintf(l.Urls.Read, ns, rc)
	return l.patch(url, []byte(data))
}
