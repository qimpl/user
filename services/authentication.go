package services

import (
	"os"

	"github.com/qimpl/authentication/db"
	"github.com/qimpl/authentication/models"

	"github.com/dgrijalva/jwt-go"
)

// CreateJwtToken return a new JSON Web Token string of a given user.
func CreateJwtToken(user *models.User) *models.TokenHash {
	claims := &models.Token{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsOwner:   user.IsOwner,
		IsAdmin:   user.IsAdmin,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))

	return &models.TokenHash{Token: tokenString}
}

// Login check user credentials and generate a new Json Web Token if they are valid.
func Login(email string, password string) (*models.TokenHash, error) {
	user, err := db.Login(email, password)
	if err != nil {
		return nil, err
	}

	return CreateJwtToken(user), nil
}
