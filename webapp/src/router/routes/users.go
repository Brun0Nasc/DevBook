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
	{
		URI:         "/search-users",
		Method:      http.MethodGet,
		Function:    controllers.LoadUsersPage,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.LoadUserProfile,
		RequestAuth: true,
	},
}
