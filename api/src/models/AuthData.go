package models

// AuthData contains user's token and id
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
