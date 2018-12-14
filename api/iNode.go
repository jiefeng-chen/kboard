package api

import (
	"kboard/template"
	"kboard/config"
	"net/http"
	"kboard/k8s"
)

type INode struct {
	Api
}

func NewINode(config *config.Config, w http.ResponseWriter, r *http.Request) *INode {
	return &INode{
		Api{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "index",
			Actions: map[string]func(){},
			R: r,
			W: w,
		},
	}
}

func (this *INode) Index() {
	// 获取node状态
	node := this.GetString("name")
	lib := k8s.NewNode(this.Config)
	if node == ""{
		this.TplEngine.Response(101, "", "缺少节点名称")
		return
	}
	data, httpErr := lib.Read(node)
	if httpErr.Code == 200 {
		this.TplEngine.Response(100, data, httpErr.Message)
		return
	}
	this.TplEngine.Response(99, "", httpErr.Message)
}

