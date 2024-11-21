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

// Get returns all posts from followed users and the user itself
func (r Posts) Get(userID uint64) ([]models.Post, error) {
	rows, err := r.db.Query(`
	SELECT DISTINCT p.*, 
		   u.nickname 
	FROM posts p 
	INNER JOIN users u ON p.author_id = u.id 
	INNER JOIN followers f ON f.user_id = u.id 
	WHERE f.follower_id = ? OR u.id = ?
	ORDER BY 1 DESC;`, userID, userID)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetByID returns a post by its ID
func (r Posts) GetByID(postID uint64) (models.Post, error) {
	rows, err := r.db.Query(`
		SELECT p.*, u.nickname FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = ?`,
		postID)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// Update updates a post in the database
func (r Posts) Update(postID uint64, post models.Post) error {
	stmt, err := r.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

// Delete deletes a post from the database
func (r Posts) Delete(postID uint64) error {
	stmt, err := r.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil{
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

// GetByUser returns all posts from a user
func (r Posts) GetByUser(userID uint64) ([]models.Post, error) {
	stmt, err := r.db.Prepare(`
	SELECT p.*, u.nickname
	FROM posts p
	INNER JOIN users u ON u.id = p.author_id
	WHERE p.author_id = ?
	ORDER BY 1 DESC`)
	if err != nil {
		return []models.Post{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Like increments the likes of a post
func (r Posts) Like(postID uint64) error {
	stmt, err := r.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}