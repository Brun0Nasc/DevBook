package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// DoLogout handles the logout process
func DoLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
