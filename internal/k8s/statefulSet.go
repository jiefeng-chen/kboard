package k8s

import (
	"encoding/json"
	"fmt"
	"kboard/config"
	"kboard/internal"
)

type IStatefulSet interface {
	internal.IK8sCore
}

type StatefulSet struct {
	internal.K8sCore
}

func NewStatefulSet(Config config.IConfig) IStatefulSet {
	return &StatefulSet{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_STATEFULE_SET,
			Urls: internal.Urls{
				Read:   "/apis/apps/v1/namespaces/%s/statefulsets/%s",
				Create: "/apis/apps/v1/namespaces/%s/statefulsets",
			},
		},
	}
}

func (l *StatefulSet) WriteToEtcd(ns string, name string, data []byte) *internal.HttpError {
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
	return &internal.HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (l *StatefulSet) Patch(ns string, name string, data []byte) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns, name)
	jsonData := l.patch(url, data)
	httpResult := internal.GetHttpCode(jsonData)
	err := internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	return err
}
