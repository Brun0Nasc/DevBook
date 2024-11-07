package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreatePost creates a post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtracteUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPostsRepository(db)

	post.ID, err = repo.Create(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts gets all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {

}

// GetPost gets a post
func GetPost(w http.ResponseWriter, r *http.Request) {

}

// UpdatePost updates a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
