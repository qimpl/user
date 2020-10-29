package router

import (
	"github.com/qimpl/authentication/handlers"

	"github.com/gorilla/mux"
)

func createUserRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetUserByID).
		Methods("GET")

	userRouter.
		HandleFunc("/register", handlers.CreateUser).
		Methods("POST")

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.UpdateUserByID).
		Methods("PUT")

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.DeleteUserByID).
		Methods("DELETE")

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/anonymize", handlers.AnonymizeUserByID).
		Methods("PUT")

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/reset/password", handlers.ResetPassword).
		Methods("PUT")

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/validate", handlers.ValidateUserAccount).
		Methods("PUT")

	userRouter.
		HandleFunc("/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/desactivate", handlers.DesactivateUserAccount).
		Methods("PUT")
}
