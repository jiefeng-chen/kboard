package main

import (
	"kboard/core"
	"github.com/gorilla/mux"
	"net/http"
)


var Config *core.Config


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/index/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root path"))
	})

	http.Handle("/", r)
}
