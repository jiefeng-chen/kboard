package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
	"resource"
)

type IStorageClass interface {
	IK8sCore
}

type StorageClass struct {
	K8sCore
}

func NewStorageClass(Config *config.Context) *StorageClass {
	return &StorageClass{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_STORAGE_CLASS,
			Urls: Urls{
				Read:   "/apis/storage.k8s.io/v1/storageclasses/%s",
				Create: "/apis/storage.k8s.io/v1/storageclasses",
			},
		},
	}
}

func (l *StorageClass) WriteToEtcd(ns string, data []byte) *HttpError {
	// 1. 检查是否已存在
	_, err := l.Read(ns)
	if err.Code == 404 {
		// 不存在，创建
		err := l.Create(ns, data)
		if err != nil {
			return err
		}
	} else {
		// 已存在，直接覆盖
		err := l.Replace(ns, data)
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

func (l *StorageClass) Create(ns string, data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Create, ns)
	jsonData := l.post(url, data)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	// 404-不存在 409-已存在
	return err
}

func (l *StorageClass) Replace(ns string, data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns)
	jsonData := l.put(url, data)
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

func (l *StorageClass) Read(ns string) (*simplejson.Json, *HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns)
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
	// 404-不存在 409-已存在
	return jsonData, err
}

func (l *StorageClass) Delete(name string) *HttpError {
	url := fmt.Sprintf(l.Urls.Read, name)
	jsonData := l.del(url)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
		// 404-不存在 409-已存在
	} else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	return err
}
