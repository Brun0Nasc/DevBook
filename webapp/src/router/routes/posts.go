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
	{
		URI:         "/posts/{postID}/dislike",
		Method:      http.MethodPost,
		Function:    controllers.DislikePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{postID}/edit",
		Method:      http.MethodGet,
		Function:    controllers.LoadEditPostPage,
		RequestAuth: true,
	},
}
