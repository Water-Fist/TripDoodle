package entities

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsDeleted bool      `json:"isDeleted"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
