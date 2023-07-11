package routes

import (
	"backend/api/handlers"
	"backend/pkg/post"
	"github.com/gofiber/fiber/v2"
)

//type Service interface {
//	InsertPost(post *entities.Post) (*entities.Post, error)
//	FetchPosts() (*[]presenter.Post, error)
//	UpdatePost(post *entities.Post) (*entities.Post, error)
//	RemovePost(ID string) error
//}
//
//type service struct {
//	repository Repository
//}
//
//func GetPost(c *fiber.Ctx) error {
//	id := c.Params("id")
//
//	var post model.GetPost
//	err := database.DB.Db.QueryRow("SELECT id, title, content, image_url, state, is_deleted FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content, &post.ImageUrl, &post.State, &post.IsDeleted)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			// No results, return a 404 status
//			return c.Status(404).JSON(fiber.Map{"error": "No post found with ID " + id})
//		} else {
//			// Database error, return a 500 status
//			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
//		}
//	}
//
//	return c.Status(200).JSON(post)
//}

//func savePost(c *fiber.Ctx) error {
//	post := new(model.getPost)
//	if err := c.BodyParser(post); err != nil {
//		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
//	}
//
//	// Prepare SQL statement
//	stmt, err := database.DB.Db.Prepare("INSERT INTO posts (title, content, image_url, state, is_deleted) VALUES ($1, $2, $3, $4, $5) RETURNING id")
//	if err != nil {
//		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
//	}
//	defer stmt.Close()
//
//	// Execute SQL statement
//	err = stmt.QueryRow(post.Title, post.Content, post.ImageUrl, post.State, post.IsDeleted).Scan(&post.ID)
//	if err != nil {
//		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
//	}
//
//	return c.Status(201).JSON(post)
//}

func PostRouter(app fiber.Router, service post.Service) {
	app.Get("/posts", handler.GetPosts(service))
	app.Post("/posts", handler.AddPost(service))
	app.Put("/posts", handler.UpdatePost(service))
	app.Delete("/posts", handler.RemovePost(service))
}
