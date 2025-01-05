package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger prints request info
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Auth verify if cookies exists
func Auth(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			 http.Redirect(w, r, "/login", http.StatusFound)
		}
		nextFunc(w, r)
	}
}
