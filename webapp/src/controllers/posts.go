package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreatePost calls API to register a new post on the database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
	}

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// LikePost calls API to like a post
func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.APIURL, postID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// DislikePost calls API to dislike a post
func DislikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.APIURL, postID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePost calls API to update a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	if err := r.ParseForm(); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
	}

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
