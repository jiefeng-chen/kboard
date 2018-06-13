package core

import (
	"net/http"
	//"github.com/gorilla/mux"
	"control"
)

func Call(c interface{}, w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//action := mux.Vars(r)["action"]
	c := control.CtlLogin{}
	c.Index()
	Call(c, w, r)
}






