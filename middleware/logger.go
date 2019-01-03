package middleware

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"kboard/utils"
)

// log
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if matched, _ := regexp.Match("^/assets/.*", []byte(r.RequestURI)); !matched {
			log.Println(fmt.Sprintf("%s %s %s | %s", r.Method, r.RequestURI, r.Proto, utils.GetIPAdress(r)))
		}
		next.ServeHTTP(w, r)
	})
}
