package router

import "github.com/gorilla/mux"

// Generate creates a new router and returns it
func Generate() *mux.Router {
	r := mux.NewRouter()

	return r
}
