package sight

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"server/api/presenter/response"
	"server/pkg/entities"
	"testing"
	"time"
)

func TestCreateSight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	newSight := &entities.Sight{
		Name:               "Test Sight",
		Type:               "Test Type",
		Province:           "Test Province",
		CityCountyDistrict: "Test CityCountyDistrict",
		LegalDong:          "Test LegalDong",
		Ri:                 "Test Ri",
		StreetNumber:       "Test StreetNumber",
		BuildingNumber:     "Test BuildingNumber",
		Longitude:          1.0,
		Latitude:           1.0,
		IsDeleted:          false,
		CreatedAt:          time.Date(2023, 9, 29, 2, 39, 17, 0, time.UTC),
		UpdatedAt:          time.Now(),
	}

	mock.ExpectQuery("INSERT INTO sights").WithArgs(
		newSight.Name,
		newSight.Type,
		newSight.Province,
		newSight.CityCountyDistrict,
		newSight.LegalDong,
		newSight.Ri,
		newSight.StreetNumber,
		newSight.BuildingNumber,
		newSight.Latitude,
		newSight.Longitude,
		false,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := repo.CreateSight(newSight)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}

func TestReadSight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	mock.ExpectQuery(`SELECT.*FROM\s+sights`).WillReturnRows(sqlmock.NewRows([]string{
		"id", "name", "latitude", "longitude", "type", "province", "city_county_district", "legal_dong", "ri", "street_number", "building_number",
	}).AddRow(1, "Test Sight", 37.5665, 126.9780, "Type", "Province", "City", "Dong", "Ri", "Street", "Building"))

	sights, err := repo.ReadSight()

	assert.Nil(t, err)
	assert.NotNil(t, sights)
	assert.Equal(t, 1, len(*sights))
	assert.Equal(t, "Test Sight", (*sights)[0].Name)
	assert.Equal(t, float32(37.5665), (*sights)[0].Latitude)
	assert.Equal(t, float32(126.9780), (*sights)[0].Longitude)
	assert.Equal(t, "Type", (*sights)[0].Type)
	assert.Equal(t, "Province", (*sights)[0].Province)
	assert.Equal(t, "City", (*sights)[0].CityCountyDistrict)
	assert.Equal(t, "Dong", (*sights)[0].LegalDong)
	assert.Equal(t, "Ri", (*sights)[0].Ri)
	assert.Equal(t, "Street", (*sights)[0].StreetNumber)
	assert.Equal(t, "Building", (*sights)[0].BuildingNumber)
}

func TestUpdateSight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	newSight := &entities.Sight{
		ID:                 1,
		Name:               "Test Sight",
		Type:               "Test Type",
		Province:           "Test Province",
		CityCountyDistrict: "Test CityCountyDistrict",
		LegalDong:          "Test LegalDong",
		Ri:                 "Test Ri",
		StreetNumber:       "Test StreetNumber",
		BuildingNumber:     "Test BuildingNumber",
		Longitude:          1.0,
		Latitude:           1.0,
	}

	mock.ExpectExec("UPDATE sights").WithArgs(
		newSight.Name,
		newSight.Type,
		newSight.Province,
		newSight.CityCountyDistrict,
		newSight.LegalDong,
		newSight.Ri,
		newSight.StreetNumber,
		newSight.BuildingNumber,
		newSight.Latitude,
		newSight.Longitude,
		sqlmock.AnyArg(),
		newSight.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.UpdateSight(newSight)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}

func TestDeleteSight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	sightID := "1"

	mock.ExpectExec("UPDATE").WithArgs(
		true,
		sqlmock.AnyArg(),
		sightID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteSight(sightID)
	assert.Nil(t, err)
}

func TestLoadSight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	latitude := float32(37.123456)
	longitude := float32(127.123456)

	mockSights := []response.SightLoad{
		{ID: 1, Name: "Sight A"},
		{ID: 2, Name: "Sight B"},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(mockSights[0].ID, mockSights[0].Name).
		AddRow(mockSights[1].ID, mockSights[1].Name)

	mock.ExpectQuery("SELECT").WithArgs(longitude, latitude).WillReturnRows(rows)

	result, err := repo.LoadSight(latitude, longitude)
	assert.Nil(t, err)
	assert.Equal(t, len(*result), 2)
	assert.Equal(t, (*result)[0].ID, mockSights[0].ID)
	assert.Equal(t, (*result)[0].Name, mockSights[0].Name)
	assert.Equal(t, (*result)[1].ID, mockSights[1].ID)
	assert.Equal(t, (*result)[1].Name, mockSights[1].Name)
}
