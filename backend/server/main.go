package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"server/api/routes"
	_ "server/docs"
	"server/pkg/post"
	"server/pkg/sight"
	"server/pkg/user"
	"time"
)

// @title TripDoodle API
// @version 1.0
// @description TripDoodle Server API Docs

// @host localhost:8000
// @BasePath /api/v1
func main() {
	// Connect with database
	db, err := databaseConnection()

	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}
	fmt.Println("Database connection success!")

	// 관광지 데이터 DB 적재 시에 사용
	//database.SightsInsert(db)

	postRepo := post.NewRepo(db)
	postService := post.NewService(postRepo)
	sightRepo := sight.NewRepo(db)
	sightService := sight.NewService(sightRepo)
	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.PostRouter(v1, postService)
	routes.SightRouter(v1, sightService)
	routes.UserRouter(v1, userService)

	log.Fatal(app.Listen(":8000"))
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
	dbUser = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
}

// Database settings
var (
	host     string
	port     string
	dbUser   string
	password string
	dbname   string
)

func databaseConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, dbUser, password, dbname)

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
