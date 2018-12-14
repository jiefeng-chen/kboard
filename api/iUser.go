package api

import (
	"kboard/template"
	"kboard/config"
	"net/http"
)

type IUser struct {
	Api
}

func NewIUser(config *config.Config, w http.ResponseWriter, r *http.Request) *IUser {
	return &IUser{
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

func (this *IUser) Index() {
	this.TplEngine.Response(100, "", "数据")
}

