package k8s

import (
	"fmt"
	"kboard/config"
	"kboard/k8s/resource"

	"github.com/bitly/go-simplejson"
)

type IStorageClass interface {
	WriteToEtcd(string, []byte) *HttpError
	Create([]byte) *HttpError
	Replace(string, []byte) *HttpError
	Read(string) (*simplejson.Json, *HttpError)
	Delete(string) *HttpError
}

type StorageClass struct {
	K8sCore
}

func NewStorageClass(Config config.IConfig) IStorageClass {
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

func (l *StorageClass) WriteToEtcd(name string, data []byte) *HttpError {
	// 1. 检查是否已存在
	_, err := l.Read(name)
	if err.Code == 404 {
		// 不存在，创建
		err := l.Create(data)
		if err != nil {
			return err
		}
	} else {
		// 已存在，返回错误
		return &HttpError{
			Code:    409,
			Message: "Exist",
			Status:  "Unknown",
		}
	}
	return &HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (l *StorageClass) Create(data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Create)
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

func (l *StorageClass) Replace(name string, data []byte) *HttpError {
	url := fmt.Sprintf(l.Urls.Read, name)
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

func (l *StorageClass) Read(name string) (*simplejson.Json, *HttpError) {
	url := fmt.Sprintf(l.Urls.Read, name)
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
