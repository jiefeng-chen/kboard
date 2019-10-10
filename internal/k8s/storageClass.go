package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"kboard/internal"
)

type IStorageClass interface {
	WriteToEtcd(string, []byte) *internal.HttpError
	Create([]byte) *internal.HttpError
	Replace(string, []byte) *internal.HttpError
	Read(string) (*simplejson.Json, *internal.HttpError)
	Delete(string) *internal.HttpError
}

type StorageClass struct {
	internal.K8sCore
}

func NewStorageClass(Config config.IConfig) IStorageClass {
	return &StorageClass{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_STORAGE_CLASS,
			Urls: internal.Urls{
				Read:   "/apis/storage.k8s.io/v1/storageclasses/%s",
				Create: "/apis/storage.k8s.io/v1/storageclasses",
			},
		},
	}
}

func (l *StorageClass) WriteToEtcd(name string, data []byte) *internal.HttpError {
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
		return &internal.HttpError{
			Code:    409,
			Message: "Exist",
			Status:  "Unknown",
		}
	}
	return &internal.HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (l *StorageClass) Create(data []byte) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Create)
	jsonData := l.post(url, data)
	httpResult := internal.GetHttpCode(jsonData)
	err := internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	// 404-不存在 409-已存在
	return err
}

func (l *StorageClass) Replace(name string, data []byte) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Read, name)
	jsonData := l.put(url, data)
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

func (l *StorageClass) Read(name string) (*simplejson.Json, *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.Read, name)
	jsonData := l.get(url)
	httpResult := internal.GetHttpCode(jsonData)
	err := internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	// 404-不存在 409-已存在
	return jsonData, err
}

func (l *StorageClass) Delete(name string) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Read, name)
	jsonData := l.del(url)
	httpResult := internal.GetHttpCode(jsonData)
	err := internal.GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind {
		err.Code = 200
		err.Message = "Success"
		// 404-不存在 409-已存在
	} else if httpResult.Code == 200 || httpResult.Status == internal.STATUS_SUCCESS {
		err.Code = 200
		err.Message = fmt.Sprintf("status:%s", err.Status)
	}
	return err
}
