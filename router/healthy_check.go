package router

import (
	"github.com/qimpl/authentication/handlers"

	"github.com/gorilla/mux"
)

func createHealthyRouter(router *mux.Router) {
	healthyRouter := router.PathPrefix("/healthy").Subrouter()

	healthyRouter.
		HandleFunc("", handlers.HealthyCheck).
		Methods("GET")
}
