package request

type SightRequest struct {
	Name               string  `json:"name"`
	Type               string  `json:"type"`
	Province           string  `json:"province"`
	CityCountyDistrict string  `json:"cityCountyDistrict"`
	LegalDong          string  `json:"legalDong"`
	Ri                 string  `json:"ri"`
	StreetNumber       string  `json:"streetNumber"`
	BuildingNumber     string  `json:"buildingNumber"`
	Latitude           float32 `json:"latitude"`
	Longitude          float32 `json:"longitude"`
}

type UpdateSightRequest struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Type               string  `json:"type"`
	Province           string  `json:"province"`
	CityCountyDistrict string  `json:"cityCountyDistrict"`
	LegalDong          string  `json:"legalDong"`
	Ri                 string  `json:"ri"`
	StreetNumber       string  `json:"streetNumber"`
	BuildingNumber     string  `json:"buildingNumber"`
	Latitude           float32 `json:"latitude"`
	Longitude          float32 `json:"longitude"`
}

type DeleteSightRequest struct {
	ID string `json:"id"`
}

type LoadSightRequest struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
