package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"kboard/internal"
)

type IResourceQuota interface {
	WriteToEtcd(string, []byte) *internal.HttpError
	Create(string, []byte) *internal.HttpError
	Replace(string, []byte) *internal.HttpError
	Read(string) (*simplejson.Json, *internal.HttpError)
}

type ResourceQuota struct {
	internal.K8sCore
	MetaName string
}

func NewResourceQuota(Config config.IConfig) IResourceQuota {
	return &ResourceQuota{
		MetaName: "quota-v1",
		K8sCore: internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_RESOURCE_QUOTA,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s/resourcequotas/%s",
				Create: "/api/v1/namespaces/%s/resourcequotas",
			},
		},
	}
}

func (l *ResourceQuota) WriteToEtcd(ns string, data []byte) *internal.HttpError {
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

	return &internal.HttpError{
		Code:    200,
		Message: "Success",
		Status:  "Unknown",
	}
}

func (l *ResourceQuota) Create(ns string, data []byte) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Create, ns)
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

func (l *ResourceQuota) Replace(ns string, data []byte) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns, l.MetaName)
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

func (l *ResourceQuota) Read(ns string) (*simplejson.Json, *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns, l.MetaName)
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
