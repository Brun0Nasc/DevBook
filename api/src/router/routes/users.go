package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequestAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.GetUsers,
		RequestAuth: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Function: controllers.GetUser,
		RequestAuth: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequestAuth: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequestAuth: false,
	},
}
