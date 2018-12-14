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
	//var node *resource.ResNode
	//	//node = resource.NewResNode("")
	// 获取node状态
	node := this.GetString("name")
	lib := k8s.NewNode(this.Config)
	this.TplEngine.Response(100, "", "数据")
}

