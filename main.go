package main

import (
	"log"

	_ "github.com/qimpl/APP_NAME/docs"
	"github.com/qimpl/APP_NAME/router"

	"github.com/joho/godotenv"
)

// @title APP_NAME API
// @version 0.1.0
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}

	router.CreateRouter()
}
