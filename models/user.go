package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct contain all the user data
type User struct {
	ID                      uuid.UUID                `json:"id,omitempty" pg:"id" swaggerignore:"true"`
	FirstName               string                   `json:"first_name" example:"Henri"`
	LastName                string                   `json:"last_name" example:"Martin"`
	Email                   string                   `json:"email" example:"henri.martin@qimpl.fr"`
	MobilePhoneNumber       string                   `json:"mobile_phone_number" example:"0610890978"`
	Birthdate               time.Time                `json:"birthdate" example:"1999-10-05"`
	Country                 string                   `json:"country" example:"FR"`
	State                   string                   `json:"state" example:"Haut-De-France"`
	Street                  string                   `json:"street" example:"Rue des peupliers"`
	City                    string                   `json:"city" example:"Lille"`
	Zip                     string                   `json:"zip" example:"59000"`
	IsOwner                 bool                     `json:"is_owner,omitempty" pg:",use_zero" example:"true"`
	Password                string                   `json:"password,omitempty" example:"MyPassword"`
	IsEnabled               bool                     `json:"is_enabled,omitempty" pg:",use_zero" example:"true"`
	IsAdmin                 bool                     `json:"is_admin,omitempty" pg:",use_zero" example:"true"`
	IsVerified              bool                     `json:"is_verified,omitempty" pg:",use_zero" example:"true"`
	IsDeleted               bool                     `json:"is_deleted,omitempty" pg:",use_zero" example:"true"`
	NotificationPreferences *NotificationPreferences `json:"notification_preferences,omitempty" pg:"rel:belongs-to"`
	StripeCustomerID        string                   `json:"stripe_customer_id,omitempty" pg:"stripe_customer_id" example:"cus_IOwdRp9gIlOjTD"`
	StripePaymentMethodID   string                   `json:"stripe_payment_method_id,omitempty" pg:"stripe_payment_method_id" example:"pm_1Ho8k8CMhQMU3AqAKJwPYAXj"`
	CreatedAt               time.Time                `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt               time.Time                `json:"updated_at,omitempty" swaggerignore:"true"`
	DeletedAt               time.Time                `json:"deleted_at,omitempty" swaggerignore:"true"`
}

//UserLogin contain only user credentials used to authentication
type UserLogin struct {
	Email    string `json:"email" example:"henri.martin@qimpl.fr"`
	Password string `json:"password" example:"MyPassword"`
}

// UserResetPasswordBody is used as struct for user reset password
type UserResetPasswordBody struct {
	Password string `json:"password"`
}

// NotificationPreferences contain all user notification preferences data
type NotificationPreferences struct {
	ID        uuid.UUID `json:"id" swaggerignore:"true"`
	UserID    uuid.UUID `json:"user_id" example:"cb7bc97f-45b0-4972-8edf-dc7300cc059c"`
	OnEmail   bool      `json:"on_email,omitempty" pg:",use_zero" example:"true"`
	OnSms     bool      `json:"on_sms,omitempty" pg:",use_zero" example:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
}
