package post

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"server/pkg/entities"
	"testing"
)

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	newPost := &entities.Post{
		Title:    "Test Title",
		Content:  "Test Content",
		ImageUrl: "http://example.com/image.jpg",
		State:    true,
		SightId:  1,
	}

	mock.ExpectQuery("INSERT INTO").WithArgs(newPost.Title, newPost.Content, newPost.ImageUrl, newPost.State, newPost.SightId, false, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := repo.CreatePost(newPost)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}

func TestReadPost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "image_url", "state", "sight_id"}).
		AddRow(1, "Test Title", "Test Content", "http://example.com/image.jpg", true, 1).
		AddRow(2, "Test Title 2", "Test Content 2", "http://example.com/image2.jpg", false, 2)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	result, err := repo.ReadPost()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), 2)
}

func TestUpdatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	updatePost := &entities.Post{
		ID:       1,
		Title:    "Test Title",
		Content:  "Test Content",
		ImageUrl: "http://example.com/image.jpg",
		State:    true,
	}

	// Use sqlmock.AnyArg() for the updated_at argument
	mock.ExpectExec("UPDATE").WithArgs(updatePost.Title, updatePost.Content, updatePost.ImageUrl, updatePost.State, false, sqlmock.AnyArg(), updatePost.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.UpdatePost(updatePost)
	assert.Nil(t, err)
	if result != nil { // Check if result is not nil before referencing its fields
		assert.Equal(t, result.ID, 1)
	}
}

func TestDeletePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	mock.ExpectExec("UPDATE").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeletePost("1")
	assert.Nil(t, err)
}
