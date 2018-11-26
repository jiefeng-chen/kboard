package middleware

import (
	"log"
	"net/http"
	"time"
)

type Session struct {
	tokenUsers map[string]string
}

func NewSession() *Session {
	return &Session{
		tokenUsers: make(map[string]string),
	}
}

// Initialize it somewhere
func (sess *Session) Populate() {
	sess.tokenUsers["00000000"] = "user0"
	sess.tokenUsers["aaaaaaaa"] = "userA"
	sess.tokenUsers["05f717e5"] = "randomUser"
	sess.tokenUsers["deadbeef"] = "user0"
}

func (sess *Session) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := sess.tokenUsers[token]; found {
			log.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func (sess *Session) SetCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "hello",
		Value:    "hello",
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	})
}
