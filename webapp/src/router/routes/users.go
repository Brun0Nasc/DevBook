package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/create-user",
		Method:      http.MethodGet,
		Function:    controllers.LoadRegisterUserPage,
		RequestAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequestAuth: false,
	},
}
