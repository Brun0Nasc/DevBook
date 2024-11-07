package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts represents the posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a new posts repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create creates a post in the database
func (r Posts) Create(post models.Post) (uint64, error) {
	stmt, err := r.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}
