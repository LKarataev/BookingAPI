package dao

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

type OrdersRepository struct {
	mutex  sync.Mutex
	orders map[int]Order
}

type OrdersRepositoryInterface interface {
	GetOrder(ID int) (Order, error)
	CreateOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (int, error)
}

func (repo *OrdersRepository) GetOrder(ID int) (Order, error) {
	if order, ok := repo.orders[ID]; ok {
		return order, nil
	}
	return Order{}, fmt.Errorf("Order with ID = %d not found in application memory", ID)
}

func (repo *OrdersRepository) CreateOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if HotelID == "" || RoomID == "" {
		return 0, fmt.Errorf("Unacceptable HotelID or RoomID")
	}

	newID := int(len(repo.orders) + 1)
	order := Order{
		HotelID:   HotelID,
		RoomID:    RoomID,
		UserEmail: UserEmail,
		From:      From,
		To:        To,
	}
	repo.orders[newID] = order
	return newID, nil
}

func NewOrdersRepository() *OrdersRepository {
	repo := OrdersRepository{}
	repo.orders = map[int]Order{}
	return &repo
}
