package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User represents a user in the system
type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// GetCompleteUser makes 4 requests on API to build the user
func GetCompleteUser(userID uint64, r *http.Request) (user User, err error) {
	chUser := make(chan User)
	chFollowers := make(chan []User)
	chFollowing := make(chan []User)
	chPosts := make(chan []Post)

	go GetUserData(chUser, userID, r)
	go GetFollowers(chFollowers, userID, r)
	go GetFollowing(chFollowing, userID, r)
	go GetPosts(chPosts, userID, r)

	var (
		followers []User
		following []User
		posts     []Post
	)

	for range 4 {
		select {
		case loadedUser := <-chUser:
			if loadedUser.ID == 0 {
				err = errors.New("user not found")
				return
			}

			user = loadedUser
		case loadedFollowers := <-chFollowers:
			if loadedFollowers == nil {
				err = errors.New("error loading followers")
				return
			}

			followers = loadedFollowers
		case loadedFollowing := <-chFollowing:
			if loadedFollowing == nil {
				err = errors.New("error loading following")
				return
			}
			following = loadedFollowing
		case loadedPosts := <-chPosts:
			if loadedPosts == nil {
				err = errors.New("error loading posts")
				return
			}
			posts = loadedPosts
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return
}

// GetUserData calls API and return the user's data
func GetUserData(ch chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		ch <- User{}
		return
	}

	ch <- user
}

// GetFollowers calls API to get user's followers
func GetFollowers(ch chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		ch <- nil
		return
	}

	ch <- followers
}

// GetFollowing calls API to get users that the user is following
func GetFollowing(ch chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		ch <- nil
		return
	}

	ch <- following
}

// GetPosts calls API to get user's posts
func GetPosts(ch chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		ch <- nil
		return
	}

	ch <- posts
}
