package control

import (
	"net/http"
	"kboard/core"
)

type CtlLogin struct {
	Control
}

func NewCtlLogin(config *core.Config, w http.ResponseWriter, r *http.Request) *CtlLogin {
	return &CtlLogin{
		Control{
			Config: config,
			TplEngine: core.NewTplEngine(w, r),
			Control: "login",
			Actions: map[string]func(){},
		},
	}
}

func (this *CtlLogin) Index() {
	this.TplEngine.Display("login")
}