package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logOutRout = Route{
	URI:         "/logout",
	Method:      http.MethodGet,
	Function:    controllers.DoLogout,
	RequestAuth: true,
}
