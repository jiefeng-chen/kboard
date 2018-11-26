package control

import (
	"flag"
	"github.com/gorilla/mux"
	"kboard/middleware"
	"net/http"
	"kboard/core"
)

type Router struct {
	Router *mux.Router
	Config *core.Config
}

func NewRouter(Config *core.Config) *Router {
	return &Router{
		Router: mux.NewRouter(),
		Config: Config,
	}
}

// register url
func (r *Router) InitRouter() *Router {
	var dir string
	flag.StringVar(&dir, "dir", "assets", "")
	flag.Parse()

	// static files
	r.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))

	// match url
	UrlRegister(r)

	// logs
	if r.Config.IsLog() {
		r.Router.Use(middleware.Logger)
	}

	// authentication
	if r.Config.IsAuth() {
		amw := middleware.NewAuthentication()
		amw.Populate()
		r.Router.Use(amw.Middleware)
	}

	// safe handler
	r.Router.Use(middleware.SafeHandler)

	return r
}


