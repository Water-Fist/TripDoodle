package user

import (
	"database/sql"
	"server/pkg/entities"
	"time"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser(user *entities.User) (*entities.User, error)
	GetUserByID(ID string) (*entities.User, error)
	EmailCheck(user *entities.User) (bool, error)
	NicknameCheck(user *entities.User) (bool, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(ID string) (*entities.User, error)
}

type NewRepository struct {
	Db *sql.DB
}

func NewRepo(Db *sql.DB) *NewRepository {
	return &NewRepository{
		Db: Db,
	}
}

func (r *NewRepository) CreateUser(user *entities.User) (*entities.User, error) {
	query :=
		`
		INSERT INTO user 
			(email, 
			 password, 
			 nickname, 
			 is_deleted, 
			 created_at, 
			 updated_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6) 
		RETURNING id
		`

	user.IsDeleted = false

	err := r.Db.QueryRow(query, user.Email, user.Password, user.Nickname, user.IsDeleted).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *NewRepository) ReadUser(user *entities.User) (*entities.User, error) {
	query :=
		`
		SELECT
			id,
			email,
			password,
			nickname,
			is_deleted,
			created_at,
			updated_at
		FROM
			user
		WHERE
			email = $1
		`

	err := r.Db.QueryRow(query, user.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Nickname, &user.IsDeleted, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *NewRepository) GetUserByID(ID string) (*entities.User, error) {
	query :=
		`
		SELECT
			id,
			email,
			nickname
		FROM
			user
		WHERE
			id = $1
		`

	user := &entities.User{}
	err := r.Db.QueryRow(query, ID).Scan(&user.ID, &user.Email, &user.Nickname)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *NewRepository) NicknameCheck(user *entities.User) (bool, error) {
	query :=
		`
		SELECT
		    EXISTS(
				SELECT
					1
				FROM
					user
				WHERE
					nickname = $1
		) AS exist
		`

	var exists bool
	err := r.Db.QueryRow(query, user.Nickname).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *NewRepository) EmailCheck(user *entities.User) (bool, error) {
	query :=
		`
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					user
				WHERE
					email = $1
		) AS exist
		`

	var exists bool
	err := r.Db.QueryRow(query, user.Email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *NewRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	query :=
		`
		UPDATE
			user
		SET
			email = $1,
			password = $2,
			nickname = $3,
		WHERE
			id = $4
		`

	_, err := r.Db.Exec(query, user.Email, user.Password, user.Nickname, user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *NewRepository) DeleteUser(ID string) error {
	query :=
		`
		UPDATE
			user
		SET
			is_deleted = $1
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
