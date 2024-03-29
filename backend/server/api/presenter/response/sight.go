package response

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/entities"
)

type Sight struct {
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

type SightLoad struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SightSuccessResponse(data *entities.Sight) *fiber.Map {
	sight := Sight{
		ID:                 data.ID,
		Name:               data.Name,
		Latitude:           data.Latitude,
		Longitude:          data.Longitude,
		Type:               data.Type,
		Province:           data.Province,
		CityCountyDistrict: data.CityCountyDistrict,
		LegalDong:          data.LegalDong,
		Ri:                 data.Ri,
		StreetNumber:       data.StreetNumber,
		BuildingNumber:     data.BuildingNumber,
	}
	return &fiber.Map{
		"state": true,
		"data":  sight,
		"error": nil,
	}
}

type SightSuccessResponseType struct {
	State bool        `json:"state"`
	Data  Sight       `json:"data"`
	Error interface{} `json:"error"`
}

func SightsSuccessResponse(data *[]Sight) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

type SightsSuccessResponseType struct {
	State bool        `json:"state"`
	Data  []Sight     `json:"data"`
	Error interface{} `json:"error"`
}

func SightsLoadSuccessResponse(data *[]SightLoad) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  data,
		"error": nil,
	}
}

func SightErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"state": true,
		"data":  "",
		"error": err.Error(),
	}
}

type SightErrorResponseType struct {
	State bool        `json:"state"`
	Data  Sight       `json:"data"`
	Error interface{} `json:"error"`
}
