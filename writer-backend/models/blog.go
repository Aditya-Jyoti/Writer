package models

import (
	"time"
)

// @Description Represents a blog with title, description, content, and optional publish date.
type Blog struct {
	Id          string `json:"id" example:"b1d5fca3-4b5b-46f3-9d13-cf6e2d2d4e12"`
	Title       string `json:"title" example:"Understanding Go Concurrency"`
	Description string `json:"description" example:"This post explains how goroutines and channels work in Go."`
	Content     string `json:"content" example:"Go makes it easy to work with concurrency by using goroutines and channels..."`

	// @default false
	Published bool `json:"published" example:"false"`

	// @default time.Now()
	UpdateDate time.Time `json:"update_date" example:"2025-05-04T12:30:00Z"`

	// PublishDate is the optional timestamp when the blog was or will be published
	PublishDate *time.Time `json:"publish_date,omitempty" example:"2025-05-05T09:00:00Z"`
}

// @Description Input structure for creating a blog post.
type CreateBlogInput struct {
	Title       string `json:"title" example:"Understanding Go Concurrency"`
	Description string `json:"description" example:"This post explains how goroutines and channels work in Go."`
	Content     string `json:"content" example:"Go makes it easy to work with concurrency by using goroutines and channels..."`
}

// @Description Input structure for updating a blog post.
type UpdateBlogInput struct {
	Title       string `json:"title" example:"Understanding Go Concurrency"`
	Description string `json:"description" example:"This post explains how goroutines and channels work in Go."`
	Content     string `json:"content" example:"Go makes it easy to work with concurrency by using goroutines and channels..."`
	Published   bool   `json:"published" example:"false"`
}

func NewBlog(title, description, content string) *Blog {
	return &Blog{
		Title:       title,
		Description: description,
		Content:     content,
		Published:   false,
		UpdateDate:  time.Now(),
		PublishDate: nil,
	}
}
