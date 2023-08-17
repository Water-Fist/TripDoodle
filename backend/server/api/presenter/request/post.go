package request

type PostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"imageUrl"`
	State    bool   `json:"state"`
	SightID  int    `json:"sightId"`
}

type UpdatePostRequest struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"imageUrl"`
	State    bool   `json:"state"`
}

type DeletePostRequest struct {
	ID string `json:"id"`
}
