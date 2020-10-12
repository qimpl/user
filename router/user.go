package router

import (
	"github.com/qimpl/authentication/handlers"

	"github.com/gorilla/mux"
)

func createUserRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.
		HandleFunc("/{user_id}", handlers.GetUserByID).
		Methods("GET")

	userRouter.
		HandleFunc("/register", handlers.CreateUser).
		Methods("POST")

	userRouter.
		HandleFunc("", handlers.UpdateUser).
		Methods("PUT")

	userRouter.
		HandleFunc("/{user_id}", handlers.DeleteUserByID).
		Methods("DELETE")
}
