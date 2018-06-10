package main

import (
	"core"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"time"
)

var (
	Config *core.Config
)

func init() {
	// init config
	Config = core.NewConfig().LoadConfigFile("config/conf.yaml")
}

func main() {
	r := core.NewRouter(Config).InitRouter()
	log.Println("Listen On", Config.GetAddress())
	server := http.Server{
		Addr:         Config.GetAddress(),
		Handler:      r.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// turn http/2 on
	if Config.IsHttp2() {
		http2.ConfigureServer(&server, &http2.Server{})
	}
	log.Println(Config.GetHttpVersion())

	if Config.IsHttps() {
		ca := Config.GetTSL()
		log.Fatal(server.ListenAndServeTLS(ca.Cert, ca.Key))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
