package main

import (
	"log"

	_ "github.com/qimpl/authentication/docs"
	"github.com/qimpl/authentication/router"

	"github.com/joho/godotenv"
)

// @title authentication API
// @version 0.1.0
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}

	router.CreateRouter()
}
