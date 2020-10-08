package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// CreateRouter create APP_NAME API routes
func CreateRouter() {
	router := mux.NewRouter()
	APIRouter := router.PathPrefix("/api/v1").Subrouter()

	APIRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	createHealthyRouter(APIRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
		},
	})

	handler := c.Handler(router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		log.Print(err)
	}
}
