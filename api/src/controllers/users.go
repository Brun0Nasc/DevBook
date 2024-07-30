package controllers

import "net/http"

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Creating a user"))
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Getting all users"))
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Getting a user"))
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser deletes a user 
func DeleteUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Deleting a user"))
}
