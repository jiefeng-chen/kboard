package core

import (
	"flag"
	"github.com/gorilla/mux"
	"middleware"
	"net/http"
)

type Router struct {
	Router *mux.Router
	Config *Config
}

func NewRouter(Config *Config) *Router {
	return &Router{
		Router: mux.NewRouter(),
		Config: Config,
	}
}

func (r *Router) InitRouter() *Router {
	var dir string
	flag.StringVar(&dir, "dir", "assets", "")
	flag.Parse()

	// static files
	r.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))

	// mux router
	r.Router.HandleFunc("/login", LoginHandler)


	// logs
	if r.Config.IsLog() {
		r.Router.Use(middleware.LoggerMiddleware)
	}

	// authentication
	if r.Config.IsAuth() {
		amw := middleware.NewAuthenticationMiddleware()
		amw.Populate()
		r.Router.Use(amw.Middleware)
	}

	// safe handler
	r.Router.Use(middleware.SafeHandlerMiddleware)

	return r
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("login"))
}
