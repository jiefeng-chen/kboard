package api

import (
	"kboard/template"
	"kboard/config"
	"net/http"
)

type IIndex struct {
	Api
}

func NewIIndex(config *config.Config, w http.ResponseWriter, r *http.Request) *IIndex {
	return &IIndex{
		Api{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "index",
			Actions: map[string]func(){},
		},
	}
}

func (this *IIndex) Index() {
	this.TplEngine.Response(100, "", "数据")
}

