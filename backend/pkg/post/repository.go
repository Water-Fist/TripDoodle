package post

import (
	"backend/api/presenter/response"
	"backend/pkg/entities"
	"database/sql"
	"time"
)

type Repository interface {
	CreatePost(Post *entities.Post) (*entities.Post, error)
	ReadPost() (*[]response.Post, error)
	UpdatePost(Post *entities.Post) (*entities.Post, error)
	DeletePost(ID string) error
}

type repository struct {
	Db *sql.DB
}

func NewRepo(Db *sql.DB) Repository {
	return &repository{
		Db: Db,
	}
}

func (r *repository) CreatePost(post *entities.Post) (*entities.Post, error) {
	query :=
		`
			INSERT INTO 
			    posts (title, content, image_url, state, sight_id, is_deleted, created_at, updated_at) 
			VALUES 
			    ($1, $2, $3, $4, $5, $6, $7, $8) 
			RETURNING id
		`

	post.IsDeleted = false
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	err := r.Db.QueryRow(query, post.Title, post.Content, post.ImageUrl, post.State, post.SightId, post.IsDeleted, post.CreatedAt, post.UpdatedAt).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) ReadPost() (*[]response.Post, error) {
	query :=
		`
			SELECT
    			id,
    			title,
    			content,
    			image_url, 
    			state,
				sight_id
			FROM 
			    posts 
			WHERE 
			    is_deleted = false
		`

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []response.Post
	for rows.Next() {
		var post response.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.ImageUrl, &post.State, &post.SightID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *repository) UpdatePost(post *entities.Post) (*entities.Post, error) {
	query :=
		`
			UPDATE 
			    posts 
			SET 
			    title = $1, 
			    content = $2, 
			    image_url = $3, 
			    state = $4, 
			    is_deleted = $5, 
			    updated_at = $6 
			WHERE 
			    id = $7
		`

	post.UpdatedAt = time.Now()

	_, err := r.Db.Exec(query, post.Title, post.Content, post.ImageUrl, post.State, post.IsDeleted, post.UpdatedAt, post.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) DeletePost(ID string) error {
	//query := `DELETE FROM posts WHERE id = $1`

	// 실제 데이터 삭제가 아닌 is_deleted를 true로 변경
	query :=
		`
			UPDATE 
			    posts 
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
