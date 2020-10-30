package services

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/qimpl/authentication/db"
	"github.com/qimpl/authentication/models"

	"github.com/dgrijalva/jwt-go"
)

// CreateJwtToken return a new JSON Web Token string of a given user.
func CreateJwtToken(user *models.User) *models.TokenHash {
	expirationDate, _ := strconv.Atoi(time.Now().Add(time.Hour * 2).Format("20060102150405"))
	claims := &models.Token{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsOwner:   user.IsOwner,
		IsAdmin:   user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(expirationDate),
			Issuer:    "Qimpl",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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

// ValidateJwtToken verify token authenticity and signature. Returns true/false & error
func ValidateJwtToken(tokenString string) (bool, error) {
	tokenString = strings.Trim(tokenString, " ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["HS256"])
		}
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, err
}
