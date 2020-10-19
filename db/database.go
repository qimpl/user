package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	// Autoload .env file
	_ "github.com/joho/godotenv/autoload"
)

// Db is the current database connection used in the API
var Db *pg.DB

func init() {
	Db = databaseConnection()

	if Db == nil {
		log.Printf("failed to connect to Database")
		os.Exit(1)
	}
}

func databaseConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Database: os.Getenv("DB_NAME"),
	})
}
