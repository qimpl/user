package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/qimpl/authentication/db"
	"github.com/qimpl/authentication/models"
)

// CreateTimeSlot create a new time slot inside the database
// @Summary Create a new time slot
// @Description Create a new time slot of a user
// @Tags Time Slots
// @Accept json
// @Param TimeSlot body models.TimeSlot true "TimeSlot information"
// @Produce json
// @Success 201 body models.TimeSlot TimeSlot
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /time-slots [post]
func CreateTimeSlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var timeSlot models.TimeSlot

	if err := json.NewDecoder(r.Body).Decode(&timeSlot); err != nil {
		var unprocessableEntity *models.UnprocessableEntity
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	if err := db.CreateTimeSlot(&timeSlot); err != nil {
		var badRequest *models.BadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during time slot creation"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(timeSlot)
}

// UpdateTimeSlotByID update a time slot from a given ID
// @Summary Update Time Slot
// @Description Update Time Slot by its ID
// @Tags Time Slots
// @Accept json
// @Param time_slot_id query string true "Time Slot ID"
// @Param TimeSlot body models.TimeSlot true "Time Slot information"
// @Produce json
// @Success 200 body models.TimeSlot TimeSlot
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /time-slots/{time_slot_id} [put]
func UpdateTimeSlotByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body models.TimeSlotUpdate

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		var unprocessableEntity *models.UnprocessableEntity
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	var timeSlot models.TimeSlot

	timeSlotID := uuid.MustParse(mux.Vars(r)["time_slot_id"])
	timeSlot, err := db.GetTimeSlotByID(timeSlotID)

	if err != nil {
		var notFound *models.NotFound
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(notFound.GetError("The given time slot ID doesn't exist"))

		return
	}

	timeSlot.StartTime = body.StartTime
	timeSlot.EndTime = body.EndTime

	if _, err := db.UpdateTimeSlotByID(&timeSlot); err != nil {
		var badRequest *models.BadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during time slot update"))

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timeSlot)
}

// GetTimeSlotsByUserID return all time slots of a given user
// @Summary Get all time slots of a user
// @Description Get all time slots of a user using his ID
// @Tags Time Slots
// @Param user_id query string true "User ID"
// @Produce json
// @Success 200 body []models.TimeSlot TimeSlot
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Router /time-slots/user/{user_id} [get]
func GetTimeSlotsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := uuid.MustParse(mux.Vars(r)["user_id"])

	if _, err := db.GetUserByID(userID); err != nil {
		var notFound *models.NotFound
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(notFound.GetError("The given user ID doesn't exist"))

		return
	}

	timeSlots, err := db.GetTimeSlotsByUserID(userID)

	if err != nil {
		var badRequest *models.BadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user time slots retrieval"))

		return
	}

	json.NewEncoder(w).Encode(timeSlots)
}
