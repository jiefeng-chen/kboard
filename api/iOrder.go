package api

import (
	"kboard/template"
	"kboard/config"
	"net/http"
)

type IOrder struct {
	Api
}

func NewIOrder(config *config.Config, w http.ResponseWriter, r *http.Request) *IOrder {
	return &IOrder{
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

func (this *IOrder) Index() {

	this.TplEngine.Response(100, "", "")
}


