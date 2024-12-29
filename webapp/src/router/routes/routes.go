package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents the routes of the application
type Route struct {
	URI         string
	Method      string
	Function    func(w http.ResponseWriter, r *http.Request)
	RequestAuth bool
}

// Configure configures the routes of the application
func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, mainPageRoute)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
