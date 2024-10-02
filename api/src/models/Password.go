package models

// Password represents the format of updating a user's password request
type Password struct {
	New     string `json:"new,omitempty"`
	Current string `json:"current,omitempty"`
}
