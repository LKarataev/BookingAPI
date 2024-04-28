package dao

import (
	"fmt"
	"time"
)

type Order struct {
	ID        int
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

type OrdersRepository struct {
	Orders []Order
}

type OrdersRepositoryInterface interface {
	GetOrder(ID int) (Order, error)
	CreateOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (int, error)
}

func (or *OrdersRepository) GetOrder(ID int) (Order, error) {
	for _, order := range or.Orders {
		if ID == order.ID {
			return order, nil
		}
	}
	return Order{}, fmt.Errorf("Order with ID=%d not found in application memory", ID)
}

func (or *OrdersRepository) CreateOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (int, error) {
	newID := int(len(or.Orders) + 1)
	order := Order{
		ID:        newID,
		HotelID:   HotelID,
		RoomID:    RoomID,
		UserEmail: UserEmail,
		From:      From,
		To:        To,
	}
	or.Orders = append(or.Orders, order)
	return newID, nil
}
