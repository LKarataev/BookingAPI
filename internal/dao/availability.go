package dao

import (
	"fmt"
	"sync"
	"time"

	"github.com/LKarataev/BookingAPI/internal/utils"
)

type RoomAvailability struct {
	HotelID string
	RoomID  string
	Date    time.Time
}

type RoomAvailabilityRepository struct {
	mutex        sync.Mutex
	availability map[RoomAvailability]int
}

type RoomAvailabilityRepositoryInterface interface {
	GetQuota(HotelID string, RoomID string, Date time.Time) (int, error)
	DecrementQuota(HotelID string, RoomID string, Date time.Time) error
	SetPreparedData()
}

func (repo *RoomAvailabilityRepository) SetPreparedData() {
	repo.availability = map[RoomAvailability]int{
		RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    utils.Date(2024, 1, 1),
		}: 0,
		RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    utils.Date(2024, 1, 2),
		}: 1,
		RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    utils.Date(2024, 1, 3),
		}: 1,
		RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    utils.Date(2024, 1, 4),
		}: 1,
		RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    utils.Date(2024, 1, 5),
		}: 0,
	}
}

func (repo *RoomAvailabilityRepository) GetQuota(HotelID string, RoomID string, Date time.Time) (int, error) {
	room := RoomAvailability{
		HotelID: HotelID,
		RoomID:  RoomID,
		Date:    Date,
	}
	if quota, ok := repo.availability[room]; ok {
		return quota, nil
	}
	return 0, fmt.Errorf("Data with HotelID = %s , RoomID = %s, Date = %s not found in application memory", HotelID, RoomID, Date.Format("2006/01/02"))
}

func (repo *RoomAvailabilityRepository) DecrementQuota(HotelID string, RoomID string, Date time.Time) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	room := RoomAvailability{
		HotelID: HotelID,
		RoomID:  RoomID,
		Date:    Date,
	}
	if quota, ok := repo.availability[room]; ok {
		quota--
		repo.availability[room] = quota
		return nil
	}
	return fmt.Errorf("Data with HotelID = %s, RoomID = %s, Date = %s not found in application memory", HotelID, RoomID, Date.Format("2006/01/02"))
}

func NewRoomAvailabilityRepository() *RoomAvailabilityRepository {
	repo := RoomAvailabilityRepository{}
	repo.availability = map[RoomAvailability]int{}
	return &repo
}
