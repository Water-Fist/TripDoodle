package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

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

type Dbinstance struct {
	Db *sql.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")

	DB = Dbinstance{
		Db: db,
	}
}
