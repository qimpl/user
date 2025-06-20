package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/qimpl/authentication/models"
	"github.com/qimpl/authentication/services"
)

// Authenticate return a JSON Web Token from a given user
// @Summary Can allow user to login and get an JWT.
// @Description Control user credential and generate a new Json Web Token
// @Tags Authentication
// @Param UserLogin body models.UserLogin true "UserLogin object"
// @Produce json
// @Success 200 body {string} string
// @Failure 400 body models.ErrorResponse
// @Failure 422 body models.ErrorResponse
// @Router /authenticate [post]
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogin
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))
		log.Printf("Authentication - Authenticate - Unprocessable Entity - %s - %s \n", time.Now(), err)
		return
	}

	tokenHash, loginErr := services.Login(userLogin.Email, userLogin.Password)
	if loginErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during user authentication"))
		log.Printf("Authentication - Authenticate - Bad Request - %s - %s \n", time.Now(), loginErr)
		return
	}

	json.NewEncoder(w).Encode(tokenHash)
}
