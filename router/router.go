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

var unProtectedRoutes = []string{"/user/register", "/user/{user_id}/partial", "/authenticate", "/swagger/"}

// CreateRouter create authentication API routes
func CreateRouter() {
	router := mux.NewRouter()
	APIRouter := router.PathPrefix("/v1").Subrouter()

	APIRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	if os.Getenv("ENV") != "dev" {
		router.Use(jwtTokenVerificationMiddleware)
	}

	createAuthenticationRouter(APIRouter)
	createUserRouter(APIRouter)
	createTimeSlotsRouter(APIRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{
			"Authorization",
		},
	})

	handler := c.Handler(router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		log.Print(err)
	}
}
