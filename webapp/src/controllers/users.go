package controllers

import (
	"fmt"
	"net/http"
)

// CreateUser calls API to register a user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")

	fmt.Println(name)
}
