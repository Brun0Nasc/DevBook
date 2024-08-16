package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger prints requisition informations on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verify if the user who is making the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validating...")
		next(w, r)
	}
}
