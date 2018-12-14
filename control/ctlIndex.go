package control


import (
	"net/http"
	"kboard/config"
	"kboard/template"
)

type CtlIndex struct {
	Control
}

func NewCtlIndex(config *config.Config, w http.ResponseWriter, r *http.Request) *CtlIndex {
	return &CtlIndex{
		Control{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "index",
			Actions: map[string]func(){},
			R: r,
			W: w,
		},
	}
}

func (this *CtlIndex) Index() {
	this.TplEngine.Display("index")
}