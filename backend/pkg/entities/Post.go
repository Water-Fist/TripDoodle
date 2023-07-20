package entities

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"imageUrl"`
	State     bool      `json:"state"`
	IsDeleted bool      `json:"isDeleted"`
	SightId   int       `json:"sightId"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
