package models

import (
	"time"
)

type Blog struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Published   bool      `json:"published"`
	PublishDate time.Time `json:"publish_date"`
	UpdateDate  time.Time `json:"update_date"`
}
