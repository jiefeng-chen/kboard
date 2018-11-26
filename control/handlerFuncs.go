package control

import (
	"net/http"
	"github.com/gorilla/mux"
	"kboard/core"
)

// 注册路由
func UrlRegister(r *Router) {
	r.Router.HandleFunc("/login/{action:[a-z]+}", LoginHandler(r.Config))

	r.Router.HandleFunc("/index/{action:[a-z]+}", IndexHandler(r.Config))

}

func LoginHandler(c *core.Config) (f func(http.ResponseWriter, *http.Request)) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		action := mux.Vars(r)["action"]
		c := NewCtlLogin(c, w, r)
		c.Register("index", c.Index).Run(action)
	}

	return handler
}

func IndexHandler(c *core.Config) (f func(http.ResponseWriter, *http.Request)) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		action := mux.Vars(r)["action"]
		c := NewCtlIndex(c, w, r)
		c.Register("index", c.Index).Run(action)
	}

	return handler
}






