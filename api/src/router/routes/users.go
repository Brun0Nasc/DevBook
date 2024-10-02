package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequestAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.GetUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}/followers",
		Method:      http.MethodGet,
		Function:    controllers.GetFollowers,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}/following",
		Method:      http.MethodGet,
		Function:    controllers.GetFollowing,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}/update-password",
		Method:      http.MethodPost,
		Function:    controllers.UpdatePassword,
		RequestAuth: true,
	},
}
