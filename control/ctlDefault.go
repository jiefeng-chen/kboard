package control

import (
	"kboard/config"
	"kboard/template"
	"net/http"
)

type CtlDefault struct {
	Control
}

func NewCtlDefault(config *config.Config, w http.ResponseWriter, r *http.Request) *CtlDefault {
	return &CtlDefault{
		Control{
			Config:    config,
			TplEngine: template.NewTplEngine(w, r),
			Module:    "default",
			Actions:   map[string]func(){},
			R:         r,
			W:         w,
		},
	}
}

func (this *CtlDefault) Index() {
	this.Response(100, "", "ok")
}
