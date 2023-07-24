package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type sightData struct {
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	Province           string    `json:"province"`
	CityCountyDistrict string    `json:"cityCountyDistrict"`
	LegalDong          string    `json:"legalDong"`
	Ri                 string    `json:"ri"`
	StreetNumber       string    `json:"streetNumber"`
	BuildingNumber     string    `json:"buildingNumber"`
	Latitude           string    `json:"latitude"`
	Longitude          string    `json:"longitude"`
	IsDeleted          bool      `json:"isDeleted"`
	DeletedAt          time.Time `json:"deletedAt"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

func SightsInsert(Db *sql.DB) {
	csvFile, err := os.Open("database/datafile/KC_495_LLR_ATRCTN_2022.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		sight := sightData{
			Name:               line[0],
			Type:               line[1],
			Province:           line[2],
			CityCountyDistrict: line[3],
			LegalDong:          line[4],
			Ri:                 line[5],
			StreetNumber:       line[6],
			BuildingNumber:     line[7],
			Longitude:          line[8],
			Latitude:           line[9],
			IsDeleted:          false,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		query :=
			`
			INSERT INTO 
			    sights (name, type, province, city_county_district, legal_dong, ri, street_number, building_number, latitude, longitude, is_deleted, created_at, updated_at) 
			VALUES 
			    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
			`

		_, err := Db.Exec(query, sight.Name, sight.Type, sight.Province, sight.CityCountyDistrict, sight.LegalDong, sight.Ri, sight.StreetNumber, sight.BuildingNumber, sight.Latitude, sight.Longitude, sight.IsDeleted, sight.CreatedAt, sight.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
	}
}
