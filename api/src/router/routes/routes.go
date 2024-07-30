package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct that represents a route in the application
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

// Configure receives a router and configures all routes
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
