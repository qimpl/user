package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct contain all the user data
type User struct {
	ID                          uuid.UUID                `json:"id,omitempty" pg:"id" swaggerignore:"true"`
	Civility                    string                   `json:"civility" example:"Mme"`
	FirstName                   string                   `json:"first_name" example:"Henri"`
	LastName                    string                   `json:"last_name" example:"Martin"`
	Email                       string                   `json:"email" example:"henri.martin@qimpl.fr"`
	MobilePhoneNumber           string                   `json:"mobile_phone_number" example:"0610890978"`
	Birthdate                   time.Time                `json:"birthdate" example:"1999-10-05T00:00:00Z"`
	Country                     string                   `json:"country" example:"FR"`
	State                       string                   `json:"state" example:"Haut-De-France"`
	Street                      string                   `json:"street" example:"Rue des peupliers"`
	AdditionalStreetInformation string                   `json:"additional_street_information,omitempty" example:"Residence des peupliers"`
	City                        string                   `json:"city" example:"Lille"`
	Zip                         string                   `json:"zip" example:"59000"`
	IsOwner                     bool                     `json:"is_owner,omitempty" pg:",use_zero" example:"true"`
	Password                    string                   `json:"password,omitempty" example:"MyPassword"`
	IsEnabled                   bool                     `json:"is_enabled,omitempty" pg:",use_zero" example:"true"`
	IsAdmin                     bool                     `json:"is_admin,omitempty" pg:",use_zero" example:"true"`
	IsDeleted                   bool                     `json:"is_deleted,omitempty" pg:",use_zero" example:"true"`
	NotificationPreferences     *NotificationPreferences `json:"notification_preferences,omitempty" pg:"rel:belongs-to"`
	UserVerifications           *UserVerifications       `json:"user_verifications,omitempty" pg:"rel:belongs-to"`
	StripeCustomerID            string                   `json:"stripe_customer_id,omitempty" pg:"stripe_customer_id" example:"cus_IOwdRp9gIlOjTD"`
	StripeAccountID             string                   `json:"stripe_account_id,omitempty" pg:"stripe_account_id" example:"acct_1HqMQH2Hlu9RYi7N"`
	StripePaymentMethodID       string                   `json:"stripe_payment_method_id,omitempty" pg:"stripe_payment_method_id" example:"pm_1Ho8k8CMhQMU3AqAKJwPYAXj"`
	CreatedAt                   time.Time                `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt                   time.Time                `json:"updated_at,omitempty" swaggerignore:"true"`
	DeletedAt                   time.Time                `json:"deleted_at,omitempty" swaggerignore:"true"`
}

// PartialUser contains only non sensible user data return inside a non protected get user API route
type PartialUser struct {
	FirstName  string    `json:"first_name" example:"Henri"`
	LastName   string    `json:"last_name" example:"Martin"`
	IsVerified bool      `json:"is_verified" pg:",use_zero" example:"true"`
	CreatedAt  time.Time `json:"created_at" example:"1977-04-22T06:00:00Z"`
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

// StripeIdentityVerificationResponse contain all data from Stripe identity API response
type StripeIdentityVerificationResponse struct {
	ID         string `json:"id"`
	NextAction struct {
		RedirectToURL string `json:"redirect_to_url"`
		Type          string `json:"type"`
	} `json:"next_action"`
	PersonID   string `json:"person"`
	ReturnURL  string `json:"return_url"`
	RefreshURL string `json:"refresh_url"`
	Status     string `json:"status"`
}

// UserVerifications contain all user verification data
type UserVerifications struct {
	ID               uuid.UUID `json:"id" swaggerignore:"true"`
	UserID           uuid.UUID `json:"user_id" example:"cb7bc97f-45b0-4972-8edf-dc7300cc059c"`
	IsVerified       bool      `json:"is_verified,omitempty" pg:",use_zero" example:"true"`
	StripePersonID   string    `json:"stripe_person_id,omitempty" example:"vip_IRVyOgajqmYpMq"`
	VerificationType string    `json:"verification_type" example:"identity_document"`
	VerifiedAt       time.Time `json:"verified_at,omitempty" example:""`
	CreatedAt        time.Time `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
}
