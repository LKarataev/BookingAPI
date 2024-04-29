package dao

import (
	"fmt"
	"time"
)

type RoomAvailability struct {
	HotelID string
	RoomID  string
	Date    time.Time
	Quota   int
}

type RoomAvailabilityRepository struct {
	availability []RoomAvailability
}

type RoomAvailabilityRepositoryInterface interface {
	GetQuota(HotelID string, RoomID string, Date time.Time) (int, error)
	DecrementQuota(HotelID string, RoomID string, Date time.Time) error
	SetPreparedData()
}

func (r *RoomAvailabilityRepository) SetPreparedData() {
	r.availability = []RoomAvailability{
		{"reddison", "lux", date(2024, 1, 1), 1},
		{"reddison", "lux", date(2024, 1, 2), 1},
		{"reddison", "lux", date(2024, 1, 3), 1},
		{"reddison", "lux", date(2024, 1, 4), 1},
		{"reddison", "lux", date(2024, 1, 5), 0},
	}
}

func (r *RoomAvailabilityRepository) GetQuota(HotelID string, RoomID string, Date time.Time) (int, error) {
	for _, a := range r.availability {
		if a.HotelID == HotelID && a.RoomID == RoomID && a.Date.Equal(Date) {
			return a.Quota, nil
		}
	}
	return 0, fmt.Errorf("Data with HotelID=%d, RoomID=%d, Date=%T not found in application memory", HotelID, RoomID, Date)
}

func (r *RoomAvailabilityRepository) DecrementQuota(HotelID string, RoomID string, Date time.Time) error {
	for i, a := range r.availability {
		if a.HotelID == HotelID && a.RoomID == RoomID && a.Date.Equal(Date) {
			r.availability[i].Quota--
			return nil
		}
	}
	return fmt.Errorf("Data with HotelID=%d, RoomID=%d, Date=%T not found in application memory", HotelID, RoomID, Date)
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
