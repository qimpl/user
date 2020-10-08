package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct contain all the user data
type User struct {
	ID           uuid.UUID `json:"id,omitempty" pg:"id" swaggerignore:"true"`
	FirstName    string    `json:"first_name" example:"Henri"`
	LastName     string    `json:"last_name" example:"Martin"`
	Email        string    `json:"email" example:"henri.martin@qimpl.fr"`
	Birthdate    time.Time `json:"birthdate" example:"1999-10-05"`
	Country      string    `json:"country" example:"FR"`
	State        string    `json:"state" example:"Haut-De-France"`
	StreetNumber string    `json:"street_number" example:"1"`
	Street       string    `json:"street" example:"Rue des peupliers"`
	City         string    `json:"city" example:"Lille"`
	Zip          string    `json:"zip" example:"59000"`
	IsOwner      bool      `json:"is_owner,omitempty" example:"true"`
	Password     string    `json:"password,omitempty" example:"MyPassword"`
	CreatedAt    time.Time `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
}

//UserLogin contain only user credentials used to authentication
type UserLogin struct {
	Email    string `json:"email" example:"henri.martin@qimpl.fr"`
	Password string `json:"password" example:"MyPassword"`
}
