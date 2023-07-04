package main

import (
	"backend/api/routes"
	"backend/pkg/post"
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {
	// Connect with database
	db, err := databaseConnection()

	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}
	fmt.Println("Database connection success!")

	postRepo := post.NewRepo(db)
	postService := post.NewService(postRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})
	api := app.Group("/api")
	routes.PostRouter(api, postService)

	log.Fatal(app.Listen(":3000"))
}

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Assign environment variables
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
}

// Database settings
var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

func databaseConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
