package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadRegisterUserPage loads the register user page
func LoadRegisterUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

// LoadMainPage loads the main page with the posts
func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)

	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

// LoadUpdatePostPage loads the page to update a post
func LoadUpdatePostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "update-post.html", post)
}

// LoadUserProfilePage loads the user list page
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("user")))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)

	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

// LoadUserProfile loads the page of a user's profile
func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	user, _ := models.GetCompleteUser(userID, r)
	fmt.Println("User:", user)
}
