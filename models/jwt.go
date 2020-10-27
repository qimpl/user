package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Token is the Json Web Token struct  used after user login.
type Token struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	IsOwner   bool      `json:"is_owner"`
	IsAdmin   bool      `json:"is_admin"`
	jwt.StandardClaims
}

// TokenHash is the hashed Json Web Token returned after authentication.
type TokenHash struct {
	Token string `json:"token"`
}
