package router

import (
	"github.com/qimpl/authentication/handlers"

	"github.com/gorilla/mux"
)

func createTimeSlotsRouter(router *mux.Router) {
	timeSlotsRouter := router.PathPrefix("/time-slots").Subrouter()

	timeSlotsRouter.
		HandleFunc("", handlers.CreateTimeSlot).
		Methods("POST")

	timeSlotsRouter.
		HandleFunc("/user/{user_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetTimeSlotsByUserID).
		Methods("GET")
}
