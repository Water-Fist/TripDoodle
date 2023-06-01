package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model

	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"imageUrl"`
	State     bool      `json:"state"`
	IsDeleted bool      `json:"isDeleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
