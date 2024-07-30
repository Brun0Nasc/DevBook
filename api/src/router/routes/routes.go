package routes

import "net/http"

// Route is a struct that represents a route in the application
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}
