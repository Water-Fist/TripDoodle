package request

type DeleteSightRequest struct {
	ID string `json:"id"`
}

type LoadSightRequest struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
