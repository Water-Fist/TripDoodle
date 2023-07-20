package sight

import (
	"backend/api/presenter"
	"backend/pkg/entities"
	"database/sql"
	"time"
)

type Repository interface {
	CreateSight(Sight *entities.Sight) (*entities.Sight, error)
	ReadSight() (*[]presenter.Sight, error)
	UpdateSight(Sight *entities.Sight) (*entities.Sight, error)
	DeleteSight(ID string) error
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
			INSERT INTO 
			    sights (name, latitude, longitude, area, is_deleted, created_at, updated_at) 
			VALUES 
			    ($1, $2, $3, $4, $5, $6, $7) 
			RETURNING id
		`

	sight.IsDeleted = false
	sight.CreatedAt = time.Now()
	sight.UpdatedAt = time.Now()

	err := r.Db.QueryRow(query, sight.Name, sight.Latitude, sight.Longitude, sight.Area, sight.IsDeleted, sight.CreatedAt, sight.UpdatedAt).Scan(&sight.ID)
	if err != nil {
		return nil, err
	}
	return sight, nil
}

func (r *repository) ReadSight() (*[]presenter.Sight, error) {
	query :=
		`
			SELECT
    			id,
				name, 
				latitude, 
				longitude,
				area
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

	var sights []presenter.Sight
	for rows.Next() {
		var sight presenter.Sight
		err := rows.Scan(&sight.ID, &sight.Name, &sight.Latitude, &sight.Longitude, &sight.Area)
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
			    latitude = $2, 
			    longitude = $3, 
			    area = $4, 
			    is_deleted = $5, 
			    updated_at = $6 
			WHERE 
			    id = $7
		`

	sight.UpdatedAt = time.Now()

	_, err := r.Db.Exec(query, sight.Name, sight.Latitude, sight.Longitude, sight.Area, sight.IsDeleted, sight.UpdatedAt, sight.ID)
	if err != nil {
		return nil, err
	}
	return sight, nil
}

func (r *repository) DeleteSight(ID string) error {
	//query := `DELETE FROM sights WHERE id = $1`

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