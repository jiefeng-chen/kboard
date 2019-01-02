package control

import (
	"net/http"
	"kboard/config"
	"kboard/template"
)

type CtlDefault struct {
	Control
}

func NewCtlDefault(config *config.Config, w http.ResponseWriter, r *http.Request) *CtlDefault {
	return &CtlDefault{
		Control{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "default",
			Actions: map[string]func(){},
			R: r,
			W: w,
		},
	}
}

func (this *CtlDefault) Index() {
	this.TplEngine.Response(100, "", "ok")
}
