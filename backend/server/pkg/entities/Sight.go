package entities

import "time"

type Sight struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	Province           string    `json:"province"`
	CityCountyDistrict string    `json:"cityCountyDistrict"`
	LegalDong          string    `json:"legalDong"`
	Ri                 string    `json:"ri"`
	StreetNumber       string    `json:"streetNumber"`
	BuildingNumber     string    `json:"buildingNumber"`
	Latitude           float32   `json:"latitude"`
	Longitude          float32   `json:"longitude"`
	IsDeleted          bool      `json:"isDeleted"`
	DeletedAt          time.Time `json:"deletedAt"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
