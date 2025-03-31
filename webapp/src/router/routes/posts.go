package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Function:    controllers.CreatePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{postID}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePost,
		RequestAuth: true,
	},
}
