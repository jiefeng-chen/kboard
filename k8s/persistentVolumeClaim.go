package k8s

import (
	"encoding/json"
	"fmt"
	"kboard/config"
	"kboard/k8s/resource"

	"gopkg.in/yaml.v2"
)

type IPersistentVolumeClaim interface {
	IK8sCore
	ReadStatus(string, string) error
	Patch(ns string, name string, data []byte) *HttpError
}

func NewPersistentVolumeClaim(Config config.IConfig) IPersistentVolumeClaim {
	return &PersistentVolumeClaim{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_PERSISTENT_VOLUME_CLAIM,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/persistentvolumeclaims/%s",
				Create: "/api/v1/namespaces/%s/persistentvolumeclaims",
			},
		},
	}
}

type PersistentVolumeClaim struct {
	K8sCore
}

func (l *PersistentVolumeClaim) ReadStatus(ns string, name string) error {
	return nil
}

func (l *PersistentVolumeClaim) WriteToEtcd(ns string, name string, data []byte) *HttpError {
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
		dataStr := resource.NewPersistentVolumeClaim()

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
	return &HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (l *PersistentVolumeClaim) Patch(ns string, name string, data []byte) *HttpError {
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
