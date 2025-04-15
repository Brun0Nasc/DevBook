package models

import (
	"encoding/json"
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
	
	return
}

// GetUserData calls API and return the user's data
func GetUserData(ch chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.PerformRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		ch  <-User{}
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
		ch  <-nil
		return
	}
	defer response.Body.Close()
}

func GetFollowing(ch chan<- []User, userID uint64, t *http.Request) {

}

func GetPosts(ch chan<- []Post, userID uint64, t *http.Request) {

}
