package model

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageUrl  string `json:"imageUrl"`
	State     bool   `json:"state"`
	IsDeleted bool   `json:"isDeleted"`
}
