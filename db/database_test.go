package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func init() {
	Db = testDatabaseConnection()

	if Db == nil {
		log.Printf("Failed to connect to Test Database")
		os.Exit(1)
	}
}

func testDatabaseConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Database: os.Getenv("DB_TEST_NAME"),
	})
}
