package routes

import (
	"api/src/controllers"
)

var routeLogin = Route{
	URI:         "/login",
	Method:      "POST",
	Function:    controllers.Login,
	RequestAuth: false,
}
