package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
	"kboard/resource"
)

type INamespace interface {
	IK8sCore
	List(string) (*simplejson.Json, error)
}

type Namespace struct {
	K8sCore
}

func NewNamespace(Config *config.Context) *Namespace {
	return &Namespace{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_POD,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s",
				Create: "/api/v1/namespaces",
			},
		},
	}
}

func (l *Namespace) WriteToEtcd(name string, data []byte) *HttpError {
	// 1. 检查是否已存在
	_, err := l.Read(name)
	if err.Code == 404 {
		// 不存在，创建
		err := l.Create(data)
		if err != nil {
			return err
		}
	} else {
		// 已存在，直接覆盖
		err := l.Replace(name, data)
		if err != nil {
			return err
		}
	}
	return &HttpError{
		Code:    200,
		Message: "Success",
		Status:  "unknown",
	}
}

func (l *Namespace) Create(data []byte) *HttpError {
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

func (l *Namespace) Replace(ns string, data []byte) *HttpError {
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

func (l *Namespace) Read(ns string) (*simplejson.Json, *HttpError) {
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

func (l *Namespace) Delete(name string) *HttpError {
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

func (l *Namespace) List(ns string) (*simplejson.Json, *HttpError) {
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
	return jsonData, err
}
