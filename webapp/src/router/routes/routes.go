package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
		if route.RequestAuth {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
