package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// CreateUser creates a user on the database
func (u *Users) CreateUser(user models.User) (id uint64, err error) {
	return
}
