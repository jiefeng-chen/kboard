package k8s

import (
	"fmt"
	"kboard/config"
	"kboard/k8s/resource"

	"github.com/bitly/go-simplejson"
)

type IReplicationController interface {
	IK8sCore
	List(string) (*simplejson.Json, *HttpError)
	Scale(string, string, int) *HttpError
	Patch(string, string, []byte) *simplejson.Json
}

type ReplicationController struct {
	K8sCore
}

func NewReplicationController(Config *config.Config) IReplicationController {
	return &ReplicationController{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_REPLICATION_CONTROLLER,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/replicationcontrollers/%s",
				Create: "/api/v1/namespaces/%s/replicationcontrollers",
			},
		},
	}
}

func (l *ReplicationController) List(ns string) (jsonData *simplejson.Json, err *HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns)
	jsonData = l.get(url)
	httpResult := GetHttpCode(jsonData)
	err = GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
		return
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
		return
	}
	return
}

func (l *ReplicationController) Scale(ns string, rc string, num int) *HttpError {
	spec := fmt.Sprintf("{\"spec\":{\"replicas\":%d}}", num)
	json := l.Patch(ns, rc, []byte(spec))
	httpResult := GetHttpCode(json)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	return err
}

func (l *ReplicationController) Patch(ns string, rc string, data []byte) *simplejson.Json {
	url := fmt.Sprintf(l.Urls.Read, ns, rc)
	return l.patch(url, []byte(data))
}
