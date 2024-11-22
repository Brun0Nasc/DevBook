package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with all routes defined
func Generate() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
