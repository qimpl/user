package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/qimpl/authentication/db"
	"github.com/qimpl/authentication/models"
	"github.com/qimpl/authentication/services"
	"github.com/qimpl/authentication/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetAllUsers return all users
// @Summary Get users from database
// @Description Get users array objects data from database
// @Tags Users
// @Produce json
// @Success 200 {object} []models.User Users
// @Failure 400 {object} models.ErrorResponse
// @Router /user [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := db.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during users retrieval"))
		log.Printf("User - GetAlluser - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetUserByID return an user from a given ID
// @Summary Get user by his id
// @Description Get user object data from database
// @Tags Users
// @Param user_id query string true "User ID"
// @Produce json
// @Success 200 {object} models.User User
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{user_id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := db.GetUserByID(uuid.MustParse(mux.Vars(r)["user_id"]))

	var badRequest *models.BadRequest
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user retrieval"))
		log.Printf("User - GetUserByID - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	user.ProfilePicture, _ = storage.GetFromBucket(fmt.Sprintf("profile_picture_%s.png", user.ID))

	json.NewEncoder(w).Encode(user)
}

// GetPartialUserByID return a partial user with a given ID
// @Summary Get partial user by its id
// @Description Get partial user object data from database
// @Tags Users
// @Param user_id query string true "User ID"
// @Produce json
// @Success 200 {object} models.PartialUser PartialUser
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{user_id}/partial [get]
func GetPartialUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userUUID := uuid.MustParse(mux.Vars(r)["user_id"])
	partialUser, err := db.GetPartialUserByID(userUUID)

	var badRequest *models.BadRequest
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during partial user retrieval"))

		log.Printf("User - GetPartialUserByID - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	partialUser.ProfilePicture, _ = storage.GetFromBucket(fmt.Sprintf("profile_picture_%s.png", userUUID))

	json.NewEncoder(w).Encode(partialUser)
}

// CreateUser create an user into database
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Param User body models.User true "User object"
// @Produce json
// @Success 201 {object} models.User User
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /user/register [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyUser models.User

	if err := json.NewDecoder(r.Body).Decode(&bodyUser); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))
		log.Printf("User - CreateUser - Unprocessable Entity - %s - %s \n", time.Now(), err)
		return
	}

	user, err := db.CreateUser(&bodyUser)

	var badRequest *models.BadRequest
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user creation"))
		log.Printf("User - CreateUser - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	if user.ProfilePicture != "" {
		decodedImage, err := base64.StdEncoding.DecodeString(user.ProfilePicture)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during profile picture decoding"))
			log.Printf("User - CreateUser - Bad Request - %s - %s \n", time.Now(), err)
			return
		}

		imageReader := bytes.NewReader(decodedImage)

		storage.AddToBucket(fmt.Sprintf("profile_picture_%s.png", user.ID), imageReader, imageReader.Size(), http.DetectContentType(decodedImage))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(services.CreateJwtToken(user))
}

// UpdateUserByID update an user from a given ID
// @Summary Update User by his ID
// @Description Update an user
// @Tags Users
// @Accept json
// @Param User body models.User true "User object"
// @Produce json
// @Success 200 {object} models.User User
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /user/{user_id} [put]
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))
		log.Printf("User - UpdateUserByID - Unprocessable Entity - %s - %s \n", time.Now(), err)
		return
	}

	user.ID = uuid.MustParse(mux.Vars(r)["user_id"])
	if err := db.UpdateUserByID(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user update"))
		log.Printf("User - UpdateUserByID - Bad Request - %s - %s \n", time.Now(), err)
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
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{user_id} [delete]
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	err := db.DeleteUserByID(uuid.MustParse(mux.Vars(r)["user_id"]))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user deletion"))
		log.Printf("User - DeleteUserByID - Bad Request - %s - %s \n", time.Now(), err)
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
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /user/{user_id}/reset/password [put]
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var userResetPassword models.UserResetPasswordBody

	if err := json.NewDecoder(r.Body).Decode(&userResetPassword); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))
		log.Printf("User - ResetPassword - Unprocessable Entity - %s - %s \n", time.Now(), err)
		return
	}

	if err := db.ResetUserPassword(uuid.MustParse(mux.Vars(r)["user_id"]), &userResetPassword); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user password reseting"))
		log.Printf("User - ResetPassword - Bad Request - %s - %s \n", time.Now(), err)
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
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{user_id}/validate [put]
func ValidateUserAccount(w http.ResponseWriter, r *http.Request) {
	if err := db.UpdateUserAccountStatus(uuid.MustParse(mux.Vars(r)["user_id"]), true); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user account activation"))
		log.Printf("User - ValidateUserAccount - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeactivateUserAccount allow user to deactivate his account.
// @Summary Deactivate user account by his id.
// @Description Update user is_enabled field into database
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 204 ""
// @Failure 400 {object} models.ErrorResponse
// @Router /user/{user_id}/deactivate [put]
func DeactivateUserAccount(w http.ResponseWriter, r *http.Request) {
	if err := db.UpdateUserAccountStatus(uuid.MustParse(mux.Vars(r)["user_id"]), false); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user account deactivation"))
		log.Printf("User - DeactivateUserAccount - Bad Request - %s - %s \n", time.Now(), err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// AnonymizeUserByID anonymize user data into the database and set is_deleted field to true.
// @Summary anonymize user account by his id.
// @Description Anonymize user data & set is_deleted boolean to true
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 204 ""
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /user/{user_id}/anonymize [put]
func AnonymizeUserByID(w http.ResponseWriter, r *http.Request) {
	user, err := db.GetUserByID(uuid.MustParse(mux.Vars(r)["user_id"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("User does not exist"))
		log.Printf("User - AnonymizeUserByID - Not Found - %s - %s \n", time.Now(), err)
		return
	}

	user = services.AnonymizeUser(user)
	user.IsDeleted = true
	user.DeletedAt = time.Now()
	if err := db.UpdateUserByID(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user update"))
		log.Printf("User - AnonymizeUserByID - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// StripeVerificationIntent generate a verification intent from Stripe
// @Summary Create Stripe link intent to validate identity for a given user
// @Description Create Stripe Link et insert into database a new identity verification process
// @Tags Users
// @Param user_id query string true "User ID"
// @Success 201 {object} models.StripeIdentityVerificationResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /user/{user_id}/identity/verification_intent [post]
func StripeVerificationIntent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := uuid.MustParse(mux.Vars(r)["user_id"])

	user, err := db.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("User does not exist"))
		log.Printf("User - StripeVerificationIntent - Not Found - %s - %s \n", time.Now(), err)
		return
	}

	var stripeVerificationType = "identity_document"

	data := url.Values{}
	data.Set("return_url", fmt.Sprintf("%s/%s", os.Getenv("STRIPE_IDENTITY_RETURN_URL"), userID.String()))
	data.Set("requested_verifications[0]", stripeVerificationType)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", "https://api.stripe.com/v1/identity/verification_intents", strings.NewReader(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Stripe-Version", "2020-08-27; identity_beta=v3")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("STRIPE_SECRET_KEY")))

	resp, err := client.Do(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during request building"))
		log.Printf("User - StripeVerificationIntent - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	var stripeIdentityVerificationResponse models.StripeIdentityVerificationResponse

	json.NewDecoder(resp.Body).Decode(&stripeIdentityVerificationResponse)

	userVerification := &models.UserVerifications{
		UserID:                     user.ID,
		StripeVerificationIntentID: stripeIdentityVerificationResponse.ID,
		Status:                     stripeIdentityVerificationResponse.Status,
		StripePersonID:             stripeIdentityVerificationResponse.PersonID,
		VerificationType:           stripeVerificationType,
	}

	if err := db.CreateOrUpdateUserVerification(userVerification); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user verification creation"))
		log.Printf("User - StripeVerificationIntent - Bad Request - %s - %s \n", time.Now(), err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(stripeIdentityVerificationResponse)
}
