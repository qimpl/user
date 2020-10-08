package router

import (
	"github.com/qimpl/APP_NAME/handlers"

	"github.com/gorilla/mux"
)

func createHealthyRouter(router *mux.Router) {
	healthyRouter := router.PathPrefix("/healthy").Subrouter()

	healthyRouter.
		HandleFunc("", handlers.HealthyCheck).
		Methods("GET")
}
