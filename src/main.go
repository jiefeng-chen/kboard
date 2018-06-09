package src

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"core"
)

var (
	Config *core.Config
)

func init() {

}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/index/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root path"))
	})

	log.Fatal("Listen On ")
}
