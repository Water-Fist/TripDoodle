package router

import (
	"backend/database"
	"backend/model"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post model.Post
	err := database.DB.Db.QueryRow("SELECT id, title, content, image_url, state, is_deleted FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content, &post.ImageUrl, &post.State, &post.IsDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			// No results, return a 404 status
			return c.Status(404).JSON(fiber.Map{"error": "No post found with ID " + id})
		} else {
			// Database error, return a 500 status
			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
		}
	}

	return c.Status(200).JSON(post)
}
