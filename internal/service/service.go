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
	ordersRepo           dao.OrdersRepositoryInterface
	roomAvailabilityRepo dao.RoomAvailabilityRepositoryInterface
	logger               logger.Logger
}

func NewBookingApi(
    ordersRepo           dao.OrdersRepositoryInterface,
    roomAvailabilityRepo dao.RoomAvailabilityRepositoryInterface,
	logger 		  		 logger.Logger,
) BookingApi {
	return BookingApi{
		ordersRepo:           ordersRepo,
		roomAvailabilityRepo: roomAvailabilityRepo,
		logger:               logger,
	}
}

func (api *BookingApi) ConfigureRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", api.orders)
	return mux
}

func (api *BookingApi) orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getOrders(w, r)
	case http.MethodPost:
		api.createOrder(w, r)
	default:
		http.Error(w, `{"error":"bad method"}`, http.StatusNotAcceptable)
	}
}

func (api *BookingApi) getOrders(w http.ResponseWriter, r *http.Request) {}

func (api *BookingApi) createOrder(w http.ResponseWriter, r *http.Request) {
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
	a1, _ := api.ordersRepo.GetOrder(1)
	a2, _ := api.ordersRepo.GetOrder(2)
	api.logger.Info("Order1: %v", a1)
	api.logger.Info("Order2: %v", a2)
}

func (api *BookingApi) SetPreparedData() {
	api.roomAvailabilityRepo.SetPreparedData()
}
