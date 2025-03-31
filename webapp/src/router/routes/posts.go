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
		URI:         "/posts/{postID}/update",
		Method:      http.MethodGet,
		Function:    controllers.LoadUpdatePostPage,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{postID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{postID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		RequestAuth: true,
	},
}
