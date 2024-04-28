package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/LKarataev/BookingAPI/internal/dao"
	"github.com/LKarataev/BookingAPI/internal/handlers"
	"github.com/LKarataev/BookingAPI/internal/logger"
)

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

type BookingApi struct {
	ordersRepo           *dao.OrdersRepository
	roomAvailabilityRepo *dao.RoomAvailabilityRepository
	logger               logger.Logger
}

func NewBookingApi(logger logger.Logger) BookingApi {
	return BookingApi{
		ordersRepo:           &dao.OrdersRepository{},
		roomAvailabilityRepo: dao.NewRoomAvailabilityRepository(),
		logger:               logger,
	}
}

func (api BookingApi) ConfigureRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", api.orders)
	return mux
}

func (api BookingApi) orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getOrders(w, r)
	case http.MethodPost:
		api.createOrder(w, r)
	default:
		http.Error(w, `{"error":"bad method"}`, http.StatusNotAcceptable)
	}
}

func (api BookingApi) getOrders(w http.ResponseWriter, r *http.Request) {}

func (api BookingApi) createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder Order
	json.NewDecoder(r.Body).Decode(&newOrder)
	req := handlers.CreateOrderRequest{
		HotelID:   newOrder.HotelID,
		RoomID:    newOrder.RoomID,
		UserEmail: newOrder.UserEmail,
		From:      newOrder.From,
		To:        newOrder.To,
	}

	ctx := context.Background()
	resp, err := handlers.NewCreateOrderHandler(api.ordersRepo, api.roomAvailabilityRepo, api.logger).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
	api.logger.Info("Order successfully created:\n\tOrder: %v\n\tID: %d", newOrder, resp.OrderID)
}
