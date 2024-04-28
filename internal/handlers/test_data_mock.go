package handlers

import (
	"fmt"
	"time"

	"github.com/LKarataev/BookingAPI/internal/dao"
)

type OrdersRepositoryMock struct {
	MockOrders []dao.Order
}

func NewOrdersRepositoryMock() *OrdersRepositoryMock {
	MockOrders := []dao.Order{
		dao.Order{
			ID:        1,
			HotelID:   "reddison",
			RoomID:    "lux",
			UserEmail: "guest@mail.ru",
			From:      time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			To:        time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC),
		},
		dao.Order{
			ID:        2,
			HotelID:   "reddison",
			RoomID:    "lux",
			UserEmail: "guest@mail.ru",
			From:      time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			To:        time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC),
		},
		dao.Order{
			ID:        3,
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
		return dao.Order{}, fmt.Errorf("GetOrder ID < 0 error")
	}
	return dao.Order{}, fmt.Errorf("GetOrder unknown error")
}

func (r OrdersRepositoryMock) CreateOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (int, error) {
	if HotelID != "reddison" || RoomID != "lux" || UserEmail != "guest@mail.ru" {
		return 0, fmt.Errorf("Data with HotelID=%s, RoomID=%s, UserEmail=%s not found in application memory", HotelID, RoomID, UserEmail)
	}
	return 4, nil
}

type RoomAvailabilityRepositoryMock struct {
	MockAvailability []dao.RoomAvailability
}

func NewRoomAvailabilityRepositoryMock() *RoomAvailabilityRepositoryMock {
	MockAvailability := []dao.RoomAvailability{
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			Quota:   1,
		},
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC),
			Quota:   1,
		},
		dao.RoomAvailability{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC),
			Quota:   1,
		},
	}
	return &RoomAvailabilityRepositoryMock{MockAvailability: MockAvailability}
}

func (repo RoomAvailabilityRepositoryMock) GetQuota(HotelID string, RoomID string, Date time.Time) (int, error) {
	if HotelID != "reddison" || RoomID != "lux" || !(Date == time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC)) {
		return 0, fmt.Errorf("Data with HotelID=%s, RoomID=%s, Date=%v not found in application memory", HotelID, RoomID, Date)
	}
	switch Date {
	case time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[0].Quota, nil
	case time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[1].Quota, nil
	case time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC):
		return repo.MockAvailability[2].Quota, nil
	}
	return 0, fmt.Errorf("Data with HotelID=%s, RoomID=%s, Date=%v not found in application memory", HotelID, RoomID, Date)
}

func (repo RoomAvailabilityRepositoryMock) DecrementQuota(HotelID string, RoomID string, Date time.Time) error {
	if HotelID != "reddison" || RoomID != "lux" || !(Date == time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 3, 0, 0, 0, 0, time.UTC) ||
		Date == time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC)) {
		return fmt.Errorf("Data with HotelID=%s, RoomID=%s, Date=%v not found in application memory", HotelID, RoomID, Date)
	}
	return nil
}
