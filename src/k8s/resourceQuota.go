package k8s

import (
	"github.com/revel/config"
	"resource"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type IResourceQuota interface {
	IK8sCore
}

type ResourceQuota struct {
	K8sCore
	MetaName string
}

func NewResourceQuota(Config *config.Context) *ResourceQuota {
	return &ResourceQuota{
		MetaName: "quota-v1",
		K8sCore: K8sCore{
			Config: Config,
			Kind: resource.RESOURCE_RESOURCE_QUOTA,
			Urls: Urls{
				Read: "/api/v1/namespaces/%s/resourcequotas/%s",
				Create: "/api/v1/namespaces/%s/resourcequotas",
			},
		},
	}
}

func (l *ResourceQuota) WriteToEtcd(ns string, data []byte) *HttpError {
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
		Code:200,
		Message: "Success",
		Status: "Unknown",
	}
}

func (l *ResourceQuota) Create(ns string, data []byte) (*HttpError) {
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

func (l *ResourceQuota) Replace(ns string, data []byte) (*HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns, l.MetaName)
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

func (l *ResourceQuota) Read(ns string) (*simplejson.Json, *HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns, l.MetaName)
	jsonData := l.get(url)
	httpResult := GetHttpCode(jsonData)
	err := GetHttpErr(httpResult)
	if httpResult.Kind == l.Kind{
		err.Code = 200
		err.Message = "Success"
	}else if httpResult.Code == 200 || httpResult.Status == STATUS_SUCCESS {
		err.Code = 200
		err.Message = httpResult.Status + ":" + httpResult.Message
	}
	// 404-不存在 409-已存在
	return jsonData, err
}




