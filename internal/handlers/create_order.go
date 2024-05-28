package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/LKarataev/BookingAPI/internal/dao"
	"github.com/LKarataev/BookingAPI/internal/logger"
	"github.com/LKarataev/BookingAPI/internal/utils"
)

type CreateOrderHandler struct {
	ordersRepo           dao.OrdersRepositoryInterface
	roomAvailabilityRepo dao.RoomAvailabilityRepositoryInterface
}

type CreateOrderResponse struct {
	OrderID int
}

type CreateOrderRequest struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

func NewCreateOrderHandler(ordersRepo dao.OrdersRepositoryInterface, roomAvailabilityRepo dao.RoomAvailabilityRepositoryInterface) CreateOrderHandler {
	return CreateOrderHandler{ordersRepo: ordersRepo, roomAvailabilityRepo: roomAvailabilityRepo}
}

func (h CreateOrderHandler) Handle(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	log := logger.GetLogger()
	log.Info("CreateOrderRequest started")

	mutex := utils.GetStringKeyLocker()
	mutex.Lock(req.HotelID)
	defer mutex.Unlock(req.HotelID)

	daysToBook := utils.DaysBetween(req.From, req.To)
	unavailableDays := make([]time.Time, 0)

	for _, dayToBook := range daysToBook {
		quota, err := h.roomAvailabilityRepo.GetQuota(req.HotelID, req.RoomID, dayToBook)
		if err != nil {
			log.Errorf("CreateOrderRequest error: %s", err)
			return CreateOrderResponse{}, err
		}
		if quota < 1 {
			unavailableDays = append(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		log.Errorf("Hotel room is not available for selected dates: %s.", utils.DatesToStr(unavailableDays))
		return CreateOrderResponse{}, fmt.Errorf("Hotel room is not available for selected dates")
	}

	OrderID, err := h.ordersRepo.CreateOrder(req.HotelID, req.RoomID, req.UserEmail, req.From, req.To)
	if err != nil {
		log.Errorf("CreateOrderRequest error: %s", err)
		return CreateOrderResponse{}, err
	}

	for _, dayToBook := range daysToBook {
		err := h.roomAvailabilityRepo.DecrementQuota(req.HotelID, req.RoomID, dayToBook)
		if err != nil {
			
			log.Errorf("CreateOrderRequest error: %s", err)
			return CreateOrderResponse{}, err
		}
	}

	resp := CreateOrderResponse{OrderID: OrderID}
	log.Info("CreateOrderRequest successed")
	return resp, nil
}
