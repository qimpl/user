package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qimpl/authentication/db"
	"github.com/qimpl/authentication/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetUserByID return an user from a given ID
// @Summary Get user by his id
// @Description Get user object data from database
// @Tags Users
// @Param user_id query string true "User ID"
// @Produce json
// @Success 200 body models.User User
// @Failure 400 body {string} string
// @Router /user/{user_id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := db.GetUserByID(uuid.MustParse(mux.Vars(r)["user_id"]))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user retrieval"))

		return
	}

	json.NewEncoder(w).Encode(user)
}

// CreateUser create an user into database
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Param User body models.User true "User object"
// @Produce json
// @Success 201 body models.User User
// @Failure 400 body {string} string
// @Failure 422 body {string} string
// @Router /user [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyUser models.User

	if err := json.NewDecoder(r.Body).Decode(&bodyUser); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	user, err := db.CreateUser(&bodyUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user creation"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser update an user
// @Summary Update User
// @Description Update an user
// @Tags Users
// @Accept json
// @Param User body models.User true "User object"
// @Produce json
// @Success 200 body models.User User
// @Failure 400 body {string} string
// @Failure 422 body {string} string
// @Router /user [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyUser models.User

	if err := json.NewDecoder(r.Body).Decode(&bodyUser); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	_, err := db.UpdateUser(&bodyUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user update"))

		return
	}

	json.NewEncoder(w).Encode(bodyUser)

}

// DeleteUserByID delete an user from a given ID
// @Summary Delete user by his id
// @Description Delete user data from database
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 204 ""
// @Failure 400 body {string} string
// @Router /user/{user_id} [delete]
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	err := db.DeleteUserByID(uuid.MustParse(mux.Vars(r)["user_id"]))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user deletion"))

		return
	}
	w.WriteHeader(http.StatusNoContent)
}
