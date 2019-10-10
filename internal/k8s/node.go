package k8s

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"kboard/internal"
)

type INode interface {
	Nodes() (*simplejson.Json, *internal.HttpError)
	Read(string) (*simplejson.Json, *internal.HttpError)
	Create([]byte) *internal.HttpError
	Replace(string, []byte) *internal.HttpError
	WriteToEtcd(string, string, []byte) *internal.HttpError
}

type Node struct {
	internal.K8sCore
}

func NewNode(Config config.IConfig) INode {
	return &Node{
		internal.K8sCore{
			Config: Config,
			Kind:   internal.RESOURCE_NODE,
			Urls: internal.Urls{
				Read:   "/api/v1/nodes/%s",
				Create: "/api/v1/nodes",
			},
		},
	}
}

func (l *Node) Read(name string) (*simplejson.Json, *internal.HttpError) {
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

func (l *Node) Create(data []byte) *internal.HttpError {
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

func (l *Node) Replace(name string, data []byte) *internal.HttpError {
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

func (l *Node) Nodes() (*simplejson.Json, *internal.HttpError) {
	jsonData := l.get(l.Urls.Create)
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
