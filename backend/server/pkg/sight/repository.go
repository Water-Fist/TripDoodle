package sight

import (
	"database/sql"
	"server/api/presenter/response"
	"server/pkg/entities"
	"time"
)

type Repository interface {
	CreateSight(Sight *entities.Sight) (*entities.Sight, error)
	ReadSight() (*[]response.Sight, error)
	UpdateSight(Sight *entities.Sight) (*entities.Sight, error)
	DeleteSight(ID string) error
	LoadSight(Latitude float32, Longitude float32) (*[]response.SightLoad, error)
}

type repository struct {
	Db *sql.DB
}

func NewRepo(Db *sql.DB) Repository {
	return &repository{
		Db: Db,
	}
}

func (r *repository) CreateSight(sight *entities.Sight) (*entities.Sight, error) {
	query :=
		`
		INSERT INTO sights 
		(name, type, province, city_county_district, legal_dong, ri, street_number, building_number, latitude, longitude, is_deleted, created_at, updated_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
		RETURNING id
   	 	`

	sight.IsDeleted = false
	sight.CreatedAt = time.Now()
	sight.UpdatedAt = time.Now()

	err := r.Db.QueryRow(query, sight.Name, sight.Type, sight.Province, sight.CityCountyDistrict, sight.LegalDong, sight.Ri, sight.StreetNumber, sight.BuildingNumber, sight.Latitude, sight.Longitude, sight.IsDeleted, sight.CreatedAt, sight.UpdatedAt).Scan(&sight.ID)
	if err != nil {
		return nil, err
	}
	return sight, nil
}

func (r *repository) ReadSight() (*[]response.Sight, error) {
	query :=
		`
			SELECT
    			id,
				name, 
				latitude, 
				longitude,
				type, 
				province, 
				city_county_district, 
				legal_dong,
				ri,
				street_number, 
				building_number
			FROM
			    sights
			WHERE 
			    is_deleted = false
		`

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sights []response.Sight
	for rows.Next() {
		var sight response.Sight
		err := rows.Scan(&sight.ID, &sight.Name, &sight.Latitude, &sight.Longitude,
			&sight.Type, &sight.Province, &sight.CityCountyDistrict, &sight.LegalDong,
			&sight.Ri, &sight.StreetNumber, &sight.BuildingNumber)
		if err != nil {
			return nil, err
		}
		sights = append(sights, sight)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &sights, nil
}

func (r *repository) UpdateSight(sight *entities.Sight) (*entities.Sight, error) {
	query :=
		`
			UPDATE 
			    sights
			SET 
				name = $1, 
				type = $2, 
				province = $3, 
				city_county_district = $4, 
				legal_dong = $5, 
				ri = $6, 
				street_number = $7, 
				building_number = $8, 
				latitude = $9, 
				longitude = $10, 
				updated_at = $11 
			WHERE id = $12
		`

	sight.UpdatedAt = time.Now()

	_, err := r.Db.Exec(
		query,
		sight.Name,
		sight.Type,
		sight.Province,
		sight.CityCountyDistrict,
		sight.LegalDong,
		sight.Ri,
		sight.StreetNumber,
		sight.BuildingNumber,
		sight.Latitude,
		sight.Longitude,
		sight.UpdatedAt,
		sight.ID,
	)

	if err != nil {
		return nil, err
	}
	return sight, nil
}

func (r *repository) DeleteSight(ID string) error {
	//query := `DELETE FROM sight WHERE id = $1`

	// 실제 데이터 삭제가 아닌 is_deleted를 true로 변경
	query :=
		`
			UPDATE 
			    sights
			SET 
			    is_deleted = $1, 
			    deleted_at = $2 
			WHERE 
			    id = $3
		`

	_, err := r.Db.Exec(query, true, time.Now(), ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) LoadSight(Latitude float32, Longitude float32) (*[]response.SightLoad, error) {
	query :=
		`
		SELECT
			id,
			name
		FROM
			sights
		WHERE
			ST_DWithin(
				geography(ST_SetSRID(ST_Point(longitude, latitude), 4326)),
				geography(ST_SetSRID(ST_Point($1, $2), 4326)),
			1000
				)
		AND
			is_deleted = false;
		`

	rows, err := r.Db.Query(query, Longitude, Latitude)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sights []response.SightLoad
	for rows.Next() {
		var sight response.SightLoad
		err := rows.Scan(&sight.ID, &sight.Name)
		if err != nil {
			return nil, err
		}
		sights = append(sights, sight)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &sights, nil
}
