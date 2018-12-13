package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"time"
	"kboard/config"
	"kboard/router"
	"kboard/exception"
)

var (
	Config *config.Config
	NotifyReloadConfig chan int
)

func init() {
	// init config
	Config = config.NewConfig().LoadConfigFile("config/conf.toml")

	// watch config file to reload
	NotifyReloadConfig = make(chan int, 1)
	go func() {
		for {
			<-NotifyReloadConfig
			Config.ReloadConfigFile()
		}
	}()

	// init db、cache、control and so on

}

func main() {
	r := router.NewRouter(Config).InitRouter()
	log.Println("Listen On", Config.GetAddress())
	server := http.Server{
		Addr:         Config.GetAddress(),
		Handler:      r.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// turn http/2.0 on
	if Config.IsHttp2() {
		err := http2.ConfigureServer(&server, &http2.Server{})
		exception.CheckError(err, 11)
	}
	log.Println(Config.GetHttpVersion())

	if Config.IsHttps() {
		ca := Config.GetTSL()
		log.Fatal(server.ListenAndServeTLS(ca.Cert, ca.Key))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
