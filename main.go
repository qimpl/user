package main

import (
	_ "github.com/qimpl/authentication/docs"
	"github.com/qimpl/authentication/router"

	_ "github.com/joho/godotenv/autoload"
)

// @title authentication API
// @version 0.1.0
// @BasePath /v1
func main() {
	router.CreateRouter()
}
