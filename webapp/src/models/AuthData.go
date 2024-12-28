package models

// AuthData contains user's id and token
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
