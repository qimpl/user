package db

import (
	"testing"

	"github.com/qimpl/authentication/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimeSlot(t *testing.T) {
	var err error

	err = CreateTimeSlot(&models.TimeSlot{
		Weekday:   "0",
		StartTime: "14:00:00",
		EndTime:   "18:00:00",
		UserID:    uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"),
	})

	assert.NoError(t, err)

	err = CreateTimeSlot(&models.TimeSlot{
		StartTime: "14:00:00",
		EndTime:   "18:00:00",
		UserID:    uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"),
	})

	assert.EqualError(t, err, "ERROR #23502 null value in column \"weekday\" violates not-null constraint")
}

func TestGetTimeSlotByID(t *testing.T) {
	var (
		timeSlot models.TimeSlot
		err      error
	)

	timeSlot, err = GetTimeSlotByID(uuid.MustParse("08760256-f183-4b07-9f28-27b954d6cbcb"))

	assert.NoError(t, err)
	assert.IsType(t, models.TimeSlot{}, timeSlot)
	assert.Equal(t, "1", timeSlot.Weekday)

	_, err = GetTimeSlotByID(uuid.MustParse("37d66db1-d834-43f1-8980-ebeec1c66e06"))

	assert.EqualError(t, err, "pg: no rows in result set")
}

func TestUpdateTimeSlotByID(t *testing.T) {
	var (
		timeSlot *models.TimeSlot
		err      error
	)

	timeSlot, err = UpdateTimeSlotByID(&models.TimeSlot{
		ID:        uuid.MustParse("08760256-f183-4b07-9f28-27b954d6cbcb"),
		Weekday:   "1",
		StartTime: "22:00:00",
		EndTime:   "23:00:00",
		UserID:    uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"),
	})

	assert.NoError(t, err)
	assert.EqualValues(t, "22:00:00", timeSlot.StartTime)
	assert.EqualValues(t, "23:00:00", timeSlot.EndTime)

	timeSlot, err = UpdateTimeSlotByID(&models.TimeSlot{
		ID:        uuid.MustParse("08760256-f183-4b07-9f28-27b954d6cbcb"),
		Weekday:   "1",
		StartTime: "22:00:00",
		UserID:    uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"),
	})

	assert.EqualError(t, err, "ERROR #23502 null value in column \"end_time\" violates not-null constraint")
}

func TestGetTimeSlotsByUserID(t *testing.T) {
	var (
		timeSlots []models.TimeSlot
		err       error
	)

	timeSlots, err = GetTimeSlotsByUserID(uuid.MustParse("8a76b8db-0490-4ff4-8bdb-dcd18b582c2f"))

	assert.NoError(t, err)
	assert.Len(t, timeSlots, 6)
	assert.EqualValues(t, "14:00:00", timeSlots[0].StartTime)

	timeSlots, _ = GetTimeSlotsByUserID(uuid.MustParse("37d66db1-d834-43f1-8980-ebeec1c66e06"))

	assert.Nil(t, timeSlots)
}
