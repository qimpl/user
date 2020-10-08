package router

import (
	"github.com/qimpl/authentication/handlers"

	"github.com/gorilla/mux"
)

func createAuthenticationRouter(router *mux.Router) {
	authenticationRouter := router.PathPrefix("/authenticate").Subrouter()

	authenticationRouter.
		HandleFunc("", handlers.Authenticate).
		Methods("POST")

}
