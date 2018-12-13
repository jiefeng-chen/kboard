package middleware

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// log
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if matched, _ := regexp.Match("^/assets/.*", []byte(r.RequestURI)); !matched {
			log.Println(fmt.Sprintf("%s %s %s [token:%s]", r.Method, r.RequestURI, r.Proto, r.Header.Get("X-Session-Token")))
		}
		next.ServeHTTP(w, r)
	})
}
