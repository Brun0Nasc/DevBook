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

// GetUser gets a user from the database by its ID
func (u *Users) GetUser(userID uint64) (user models.User, err error) {
	row := u.db.QueryRow("SELECT id, username, nickname, email, created_at FROM users WHERE id = ?", userID)

	if err = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
		return models.User{}, err
	}

	return
}

func (u *Users) UpdateUser(userID uint64, user models.User) (err error) {
	stmt, err := u.db.Prepare("UPDATE users SET username = ?, nickname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return
}

// DeleteUser deletes a user from the database by its ID
func (u *Users) DeleteUser(userID uint64) (err error) {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID); err != nil {
		return err
	}

	return
}

// FindByEmail finds a user by its email
func (u *Users) FindByEmail(email string) (user models.User, err error) {
	row := u.db.QueryRow("SELECT id, pass FROM users WHERE email = ?", email)
	if err = row.Scan(&user.ID, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with the email %s", email)
		}
		
		return models.User{}, err
	}

	return
}
