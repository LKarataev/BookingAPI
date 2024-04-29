package handlers

import (
	"fmt"
	"time"

	"github.com/LKarataev/BookingAPI/internal/dao"
)

type OrdersRepositoryMock struct {
	MockOrders map[int]dao.Order
}

func NewOrdersRepositoryMock() *OrdersRepositoryMock {
	MockOrders := map[int]dao.Order{
		1: dao.Order{
			HotelID:   "reddison",
			RoomID:    "lux",
			UserEmail: "guest@mail.ru",
			From:      time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			To:        time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC),
		},
		2: dao.Order{
			HotelID:   "reddison",
			RoomID:    "lux",
			UserEmail: "guest@mail.ru",
			From:      time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			To:        time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC),
		},
		3: dao.Order{
			HotelID:   "reddison",
			RoomID:    "lux",
			UserEmail: "guest@mail.ru",
			From:      time.Date(2023, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			To:        time.Date(2023, time.Month(1), 5, 0, 0, 0, 0, time.UTC),
		},
	}
	return &OrdersRepositoryMock{MockOrders: MockOrders}
}

func (r OrdersRepositoryMock) GetOrder(ID int) (dao.Order, error) {
	switch ID {
	case 1, 2, 3:
		return r.MockOrders[ID-1], nil
	case -1:
		return dao.Order{}, fmt.Errorf("GetOrder ID should be greater than 0")
	}
	return dao.Order{}, fmt.Errorf("GetOrder unknown error")
}

func (r OrdersRepositoryMock) CreateOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (int, error) {
	if HotelID != "reddison" || RoomID != "lux" || UserEmail != "guest@mail.ru" {
		return 0, fmt.Errorf("Data with HotelID = %s, RoomID = %s, UserEmail = %s not found in application memory", HotelID, RoomID, UserEmail)
	}
	return 4, nil
}

type RoomAvailabilityRepositoryMock struct {
	MockAvailability map[dao.RoomAvailability]int
}

func NewRoomAvailabilityRepositoryMock() *RoomAvailabilityRepositoryMock {
	MockAvailability := map[dao.RoomAvailability]int{
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
		}: 0,
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
		}: 1,
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC),
		}: 1,
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC),
		}: 1,
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC),
		}: 0,
	}
	return &RoomAvailabilityRepositoryMock{MockAvailability: MockAvailability}
}

func (repo RoomAvailabilityRepositoryMock) GetQuota(HotelID string, RoomID string, Date time.Time) (int, error) {
	if HotelID != "reddison" || RoomID != "lux" || !(Date == time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC)) {
		return 0, fmt.Errorf("Data with HotelID = %s, RoomID = %s, Date = %s not found in application memory", HotelID, RoomID, Date.Format("2006/01/02"))
	}
	r := dao.RoomAvailability{
		HotelID: HotelID,
		RoomID:  RoomID,
		Date:    Date,
	}
	switch Date {
	case time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[r], nil
	case time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[r], nil
	case time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[r], nil
	case time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[r], nil
	case time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[r], nil
	}
	return 0, fmt.Errorf("Data with HotelID = %s, RoomID = %s, Date = %s not found in application memory", HotelID, RoomID, Date.Format("2006/01/02"))
}

func (repo RoomAvailabilityRepositoryMock) DecrementQuota(HotelID string, RoomID string, Date time.Time) error {
	if HotelID != "reddison" || RoomID != "lux" || !(Date == time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC)) {
		return fmt.Errorf("Data with HotelID = %s, RoomID = %s, Date = %s not found in application memory", HotelID, RoomID, Date.Format("2006/01/02"))
	}
	return nil
}

func (repo RoomAvailabilityRepositoryMock) SetPreparedData() {}
