package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will format and validate the user data
func (user *User) Prepare(operation string) (err error){
	if err = user.validate(operation); err != nil {
		return err
	}

	user.format()
	return
}

func (user *User) validate(operation string) (err error) {
	if user.Name == "" {
		return errors.New("the field name is required")
	}

	if user.Nick == "" {
		return errors.New("the field nick is required")
	}

	if user.Email == "" {
		return errors.New("the field email is required")
	}

	if operation == "create" && user.Password == "" {
		return errors.New("the field password is required")
	}

	return
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
