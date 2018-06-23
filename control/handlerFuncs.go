package control

import (
	"net/http"
	"github.com/gorilla/mux"
	"kboard/core"
)

func UrlRegister(r *Router) {
	r.Router.HandleFunc("/login/{action:[a-z]+}", LoginHandler(r.Config))



}

func LoginHandler(c *core.Config) (f func(http.ResponseWriter, *http.Request)) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		action := mux.Vars(r)["action"]
		c := NewCtlLogin(c, w, r)
		c.Register("index", c.Index).Run(action)
	}

	return handler
}






