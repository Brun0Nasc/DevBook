package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
	stmt, err := u.db.Prepare("INSERT INTO users (username, nickname, email, pass) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	id = uint64(lastInsertedID)

	return
}

// GetUsers gets all users from the database and filters by name or nickname
func (u *Users) GetUsers(nameOrNick string) (users []models.User, err error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := u.db.Query("SELECT id, username, nickname, email, created_at FROM users WHERE username LIKE ? OR nickname LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return
}
