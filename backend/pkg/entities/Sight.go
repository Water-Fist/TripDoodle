package entities

import "time"

type Sight struct {
	ID        int       `json:"id"`
	Name      string    `json:"title"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Area      bool      `json:"area"`
	IsDeleted bool      `json:"isDeleted"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
