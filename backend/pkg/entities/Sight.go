package entities

import "time"

type Sight struct {
	ID        int       `json:"id"`
	Name      string    `json:"title"`
	Latitude  string    `json:"content"`
	Longitude string    `json:"imageUrl"`
	Area      bool      `json:"state"`
	IsDeleted bool      `json:"isDeleted"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
