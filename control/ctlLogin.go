package control

import (
	"net/http"
	"kboard/config"
	"kboard/template"
)

type CtlLogin struct {
	Control
}

func NewCtlLogin(config *config.Config, w http.ResponseWriter, r *http.Request) *CtlLogin {
	return &CtlLogin{
		Control{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "login",
			Actions: map[string]func(){},
		},
	}
}

func (this *CtlLogin) Index() {
	this.TplEngine.Display("login")
}