package models

import (
	"time"

	"github.com/google/uuid"
)

// TimeSlot represent a single Time Slot inside the database
type TimeSlot struct {
	ID        uuid.UUID `json:"id,omitempty" pg:"id" swaggerignore:"true"`
	Weekday   string    `json:"weekday" example:"mon"`
	StartTime string    `json:"start_time" example:"14:00:00"`
	EndTime   string    `json:"end_time" example:"18:00:00"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
}
