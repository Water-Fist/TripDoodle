package entities

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"imageUrl"`
	State     bool      `json:"state"`
	IsDeleted bool      `json:"isDeleted"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
