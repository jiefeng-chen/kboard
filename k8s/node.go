package k8s

import (
	"kboard/k8s/resource"
	"github.com/bitly/go-simplejson"
	"kboard/config"
	"fmt"
)

type INode interface {
	Nodes() (*simplejson.Json, *HttpError)
	Read(string) (*simplejson.Json, *HttpError)
	Create([]byte) *HttpError
	Replace(string, []byte) *HttpError
	WriteToEtcd(string, string, []byte) *HttpError
}

type Node struct {
	K8sCore
}

func NewNode(Config *config.Config) INode {
	return &Node{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_NODE,
			Urls: Urls{
				Read:   "/api/v1/nodes/%s",
				Create: "/api/v1/nodes",
			},
		},
	}
}

func (l *Node) Read(name string) (*simplejson.Json, *HttpError) {
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

func (l *Node) Create(data []byte) *HttpError {
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

func (l *Node) Replace(name string, data []byte) *HttpError {
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


func (l *Node) Nodes() (*simplejson.Json, *HttpError) {
	jsonData := l.get(l.Urls.Create)
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
