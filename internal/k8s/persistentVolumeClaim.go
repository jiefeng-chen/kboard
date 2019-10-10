package k8s

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"kboard/config"
	"kboard/internal"
)

type IPersistentVolumeClaim interface {
	internal.IK8sCore
	ReadStatus(string, string) error
	Patch(ns string, name string, data []byte) *internal.HttpError
}

func NewPersistentVolumeClaim(Config config.IConfig) IPersistentVolumeClaim {
	return &PersistentVolumeClaim{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_PERSISTENT_VOLUME_CLAIM,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/persistentvolumeclaims/%s",
				Create: "/api/v1/namespaces/%s/persistentvolumeclaims",
			},
		},
	}
}

type PersistentVolumeClaim struct {
	internal.K8sCore
}

func (l *PersistentVolumeClaim) ReadStatus(ns string, name string) error {
	return nil
}

func (l *PersistentVolumeClaim) WriteToEtcd(ns string, name string, data []byte) *internal.HttpError {
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
		dataStr := internal.NewPersistentVolumeClaim()

		_ = yaml.Unmarshal(data, dataStr)
		yamlData := map[string]interface{}{"Spec": map[string]interface{}{
			"Resources": map[string]interface{}{
				"Requests": map[string]interface{}{
					"Storage": dataStr.GetStorage(),
				},
			},
		}}
		data2, _ := json.Marshal(yamlData)
		err := l.Patch(ns, name, data2)
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

func (l *PersistentVolumeClaim) Patch(ns string, name string, data []byte) *internal.HttpError {
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
