package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

// Prepare will format and validate the post data
func (p *Post) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()

	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("the field title is required")
	}

	if p.Content == "" {
		return errors.New("the field content is required")
	}

	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
