package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Function:    controllers.CreatePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Function:    controllers.GetPosts,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{post_id}",
		Method:      http.MethodGet,
		Function:    controllers.GetPost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{post_id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{post_id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userID}/posts",
		Method:      http.MethodGet,
		Function:    controllers.GetPostsByUser,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{post_id}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{post_id}/dislike",
		Method:      http.MethodPost,
		Function:    controllers.DislikePost,
		RequestAuth: true,
	},
}
