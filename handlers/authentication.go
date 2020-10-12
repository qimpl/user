package handlers

import (
	"encoding/json"
	"net/http"

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
// @Failure 422 body {string} string
// @Router /authenticate [post]
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		http.Error(w, "Malformed body", http.StatusUnprocessableEntity)

		return
	}

	tokenHash, loginErr := services.Login(userLogin.Email, userLogin.Password)
	if loginErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occured during user authentication")

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenHash)
}
