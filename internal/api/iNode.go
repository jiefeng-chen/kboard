package api

import (
	"kboard/config"
	"net/http"
)

type INode struct {
	IApi
}

func NewINode(config config.IConfig, w http.ResponseWriter, r *http.Request) *INode {
	node := &INode{
		IApi: *NewIApi(config, w, r),
	}
	node.Module = "node"
	return node
}

func (this *INode) Index() {

	this.TplEngine.Response(99, "", "")
}

// @todo 节点扩容
func (this *INode) Scale() {

	this.TplEngine.Response(100, "", "数据")
}

// @todo 节点隔离与恢复

// @todo 节点移除
func (this *INode) Delete() {

}
