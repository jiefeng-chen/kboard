package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// log
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto))
		next.ServeHTTP(w, r)
	})
}
