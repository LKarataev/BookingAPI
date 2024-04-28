package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/LKarataev/BookingAPI/internal/dao"
	"github.com/LKarataev/BookingAPI/internal/logger"
)

type CreateOrderHandler struct {
	ordersRepo           dao.OrdersRepositoryInterface
	roomAvailabilityRepo dao.RoomAvailabilityRepositoryInterface
	logger               logger.Logger
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

func NewCreateOrderHandler(ordersRepo dao.OrdersRepositoryInterface, roomAvailabilityRepo dao.RoomAvailabilityRepositoryInterface, logger logger.Logger) CreateOrderHandler {
	return CreateOrderHandler{ordersRepo: ordersRepo, roomAvailabilityRepo: roomAvailabilityRepo, logger: logger}
}

func (h CreateOrderHandler) Handle(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	h.logger.Info("CreateOrder request started")

	daysToBook := daysBetween(req.From, req.To)

	unavailableDays := make(map[time.Time]struct{})

	for _, dayToBook := range daysToBook {
		quota, err := h.roomAvailabilityRepo.GetQuota(req.HotelID, req.RoomID, dayToBook)
		if err != nil {
			h.logger.Errorf("CreateOrderRequest error: %s", err)
			return CreateOrderResponse{}, err
		}
		if quota < 1 {
			unavailableDays[dayToBook] = struct{}{}
		}
	}

	if len(unavailableDays) != 0 {
		h.logger.Errorf("Hotel room is not available for selected dates:\n%v", unavailableDays)
		return CreateOrderResponse{}, fmt.Errorf("Hotel room is not available for selected dates")
	}

	OrderID, err := h.ordersRepo.CreateOrder(req.HotelID, req.RoomID, req.UserEmail, req.From, req.To)
	if err != nil {
		h.logger.Errorf("CreateOrderRequest error: ", err)
		return CreateOrderResponse{}, err
	}

	for _, dayToBook := range daysToBook {
		err := h.roomAvailabilityRepo.DecrementQuota(req.HotelID, req.RoomID, dayToBook)
		if err != nil {
			h.logger.Errorf("CreateOrderRequest error: ", err)
			return CreateOrderResponse{}, err
		}
	}

	resp := CreateOrderResponse{OrderID: OrderID}
	h.logger.Info("CreateOrderRequest successed")
	return resp, nil
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
