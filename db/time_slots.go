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

// GetTimeSlotByID returns a given time slot from the database
func GetTimeSlotByID(timeSlotID uuid.UUID) (models.TimeSlot, error) {
	var timeSlot models.TimeSlot

	if err := Db.Model(&timeSlot).Where("id = ?", timeSlotID).Select(); err != nil {
		return timeSlot, err
	}

	return timeSlot, nil
}

// UpdateTimeSlotByID updates a given time slot
func UpdateTimeSlotByID(timeSlot *models.TimeSlot) (*models.TimeSlot, error) {
	if _, err := Db.Model(timeSlot).Where("id = ?", timeSlot.ID).Update(); err != nil {
		return timeSlot, err
	}

	return timeSlot, nil
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
