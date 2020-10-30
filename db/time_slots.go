package db

import (
	"github.com/google/uuid"
	"github.com/qimpl/authentication/models"
)

// CreateTimeSlot add a new time slot inside the database
func CreateTimeSlot(timeSlot *models.TimeSlot) error {
	if _, err := Db.Model(timeSlot).Insert(); err != nil {
		return err
	}

	return nil
}

// GetTimeSlotsByUserID returns every time slots of a given user
func GetTimeSlotsByUserID(userID uuid.UUID) ([]models.TimeSlot, error) {
	var timeSlots []models.TimeSlot

	if err := Db.Model(&timeSlots).
		Where("user_id = ?", userID).
		Order("weekday").
		Order("start_time").
		Select(); err != nil {
		return nil, err
	}

	return timeSlots, nil
}
