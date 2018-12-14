package k8s

import (
	"kboard/k8s/resource"
	"github.com/bitly/go-simplejson"
	"github.com/revel/config"
)

type INode interface {
	IK8sCore
	Nodes() (*simplejson.Json, *HttpError)
}

type Node struct {
	K8sCore
}

func NewNode(Config *config.Context) *Node {
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
