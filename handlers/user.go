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
// @Failure 400 {string} models.ErrorResponse
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
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
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

// UpdateUserByID update an user from a given ID
// @Summary Update User by his ID
// @Description Update an user
// @Tags Users
// @Accept json
// @Param User body models.User true "User object"
// @Produce json
// @Success 200 body models.User User
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /user/{user_id} [put]
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	user.ID = uuid.MustParse(mux.Vars(r)["user_id"])
	if _, err := db.UpdateUserByID(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user update"))

		return
	}

	json.NewEncoder(w).Encode(user)

}

// DeleteUserByID delete an user from a given ID
// @Summary Delete user by his id
// @Description Delete user data from database
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
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

// ResetPassword allow user to change his password.
// @Summary Reset user password by his id
// @Description Update user password field into database
// @Tags Users
// @Param UserResetPasswordBody body models.UserResetPasswordBody true "UserResetPasswordBody object"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /user/{user_id}/reset/password [put]
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var userResetPassword models.UserResetPasswordBody

	if err := json.NewDecoder(r.Body).Decode(&userResetPassword); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	if err := db.ResetUserPassword(uuid.MustParse(mux.Vars(r)["user_id"]), &userResetPassword); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user password reseting"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ValidateUserAccount allow user to validate his account after create it.
// @Summary Validate user account by his id.
// @Description Update user is_enabled field into database
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
// @Router /user/{user_id}/validate [put]
func ValidateUserAccount(w http.ResponseWriter, r *http.Request) {
	if err := db.UpdateUserAccountStatus(uuid.MustParse(mux.Vars(r)["user_id"]), true); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user account activation"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DesactivateUserAccount allow user to desactivate his account.
// @Summary Desactivate user account by his id.
// @Description Update user is_enabled field into database
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
// @Router /user/{user_id}/desactivate [put]
func DesactivateUserAccount(w http.ResponseWriter, r *http.Request) {
	if err := db.UpdateUserAccountStatus(uuid.MustParse(mux.Vars(r)["user_id"]), false); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user account desactivation"))

		return
	}
	w.WriteHeader(http.StatusNoContent)
}
