package models

import (
	"time"
)

type Blog struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	Published   bool       `json:"published"`
	UpdateDate  time.Time  `json:"update_date"`
	PublishDate *time.Time `json:"publish_date,omitempty"`
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
