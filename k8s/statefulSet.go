package k8s


import (
	"kboard/k8s/resource"
	"kboard/config"
	"encoding/json"
	"fmt"
)

type IStatefulSet interface {
	IK8sCore
}

type StatefulSet struct {
	K8sCore
}

func NewStatefulSet(Config *config.Config) *StatefulSet {
	return &StatefulSet{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_STATEFULE_SET,
			Urls: Urls{
				Read:   "/apis/apps/v1/namespaces/%s/statefulsets/%s",
				Create: "/apis/apps/v1/namespaces/%s/statefulsets",
			},
		},
	}
}

func (l *StatefulSet) WriteToEtcd(ns string, name string, data []byte) *HttpError {
	// 1. 检查是否已存在
	_, err := l.Read(ns, name)
	if err.Code == 404 {
		// 不存在，创建
		err := l.Create(ns, data)
		if err != nil {
			return err
		}
	} else {
		// 已存在，部分更新
		data2, _ := json.Marshal(data)
		err := l.Replace(ns, name, data2)
		if err != nil {
			return err
		}
	}
	return &HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (l *StatefulSet) Patch(ns string, name string, data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns, name)
	jsonData := l.patch(url, data)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return err
}




