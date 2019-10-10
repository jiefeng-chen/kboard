package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"kboard/internal"
)

type INamespace interface {
	WriteToEtcd(string, []byte) *internal.HttpError
	Create([]byte) *internal.HttpError
	Replace(string, []byte) *internal.HttpError
	Read(string) (*simplejson.Json, *internal.HttpError)
	Delete(string) *internal.HttpError
	List() (*simplejson.Json, *internal.HttpError)
	Watch(ns string) (*simplejson.Json, *internal.HttpError)
}

type Namespace struct {
	internal.K8sCore
}

func NewNamespace(Config config.IConfig) INamespace {
	return &Namespace{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_POD,
			Urls: internal.Urls{
				Read:   "/api/v1/namespaces/%s",
				Create: "/api/v1/namespaces",
				List:   "/api/v1/namespaces",
				Watch:  "/api/v1/watch/namespaces/%s",
			},
		},
	}
}

func (l *Namespace) WriteToEtcd(name string, data []byte) *internal.HttpError {
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
	return &internal.HttpError{
		Code:    200,
		Message: "Success",
		Status:  "unknown",
	}
}

func (l *Namespace) Create(data []byte) *internal.HttpError {
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

func (l *Namespace) Replace(ns string, data []byte) *internal.HttpError {
	url := fmt.Sprintf(l.Urls.Read, ns)
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

func (l *Namespace) Read(ns string) (*simplejson.Json, *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.Read, ns)
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

func (l *Namespace) Delete(name string) *internal.HttpError {
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

func (l *Namespace) List() (*simplejson.Json, *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.List)
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
	return jsonData, err
}

func (l *Namespace) Watch(ns string) (*simplejson.Json, *internal.HttpError) {
	url := fmt.Sprintf(l.Urls.Watch, ns)
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
	return jsonData, err
}
