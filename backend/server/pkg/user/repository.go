package user

import (
	"database/sql"
	"server/pkg/entities"
	"time"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUsers() (*[]entities.User, error)
	GetUserByID(ID string) (*entities.User, error)
	CheckEmail(email string) (bool, error)
	CheckNickname(email string) (bool, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(ID string) error
	CheckUser(email string, password string) (bool, error)
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
		INSERT INTO users 
			(name,
			 email, 
			 password, 
			 nickname, 
			 is_deleted
			) 
		VALUES 
		($1, $2, $3, $4, $5) 
		RETURNING id
		`

	user.IsDeleted = false

	err := r.Db.QueryRow(query, user.Name, user.Email, user.Password, user.Nickname, user.IsDeleted).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *NewRepository) GetUsers() (*[]entities.User, error) {
	query :=
		`
		SELECT
			id,
			email,
			password,
			nickname,
			name
		FROM
			users
		WHERE
			is_deleted = $1
		`

	rows, err := r.Db.Query(query, false)
	if err != nil {
		return nil, err
	}

	var users []entities.User

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Nickname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *NewRepository) GetUserByID(ID string) (*entities.User, error) {
	query :=
		`
		SELECT
			id,
			email,
			nickname,
			password,
			name
		FROM
			users
		WHERE
			id = $1 and is_deleted = $2
		`

	user := &entities.User{}
	err := r.Db.QueryRow(query, ID, false).Scan(&user.ID, &user.Email, &user.Nickname, &user.Password, &user.Name)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *NewRepository) CheckNickname(nickname string) (bool, error) {
	query :=
		`
		SELECT
		    EXISTS(
				SELECT
					1
				FROM
					users
				WHERE
					nickname = $1 and is_deleted = $2
		) AS exist
		`

	var exists bool
	err := r.Db.QueryRow(query, nickname, false).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *NewRepository) CheckEmail(email string) (bool, error) {
	query :=
		`
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					users
				WHERE
					email = $1 and is_deleted = $2
		) AS exist
		`

	var exists bool
	err := r.Db.QueryRow(query, email, false).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *NewRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	query :=
		`
		UPDATE
			users
		SET
			email = $1,
			password = $2,
			nickname = $3
		WHERE
			id = $4
		`

	_, err := r.Db.Exec(query, user.Email, user.Password, user.Nickname, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *NewRepository) DeleteUser(ID string) error {
	query :=
		`
		UPDATE
			users
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

// TODO: 패스워드 암호화 필요
func (r *NewRepository) CheckUser(email string, password string) (bool, error) {
	query :=
		`
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					users
				WHERE
					email = $1 AND password = $2
		) AS exist
		`

	var exists bool
	err := r.Db.QueryRow(query, email, password).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
