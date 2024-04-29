package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

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
}

func NewInMemoryBookingApi() BookingApi {
	ordersRepo := dao.NewOrdersRepository()
	roomAvailabilityRepo := dao.NewRoomAvailabilityRepository()
	return BookingApi{
		ordersRepo:           ordersRepo,
		roomAvailabilityRepo: roomAvailabilityRepo,
	}
}

func (api *BookingApi) ConfigureRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/orders", api.createOrder)
	return router
}

func (api *BookingApi) createOrder(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	var newOrder Order

	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		log.Errorf("createOrder - %s", err.Error())
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.CreateOrderRequest{
		HotelID:   newOrder.HotelID,
		RoomID:    newOrder.RoomID,
		UserEmail: newOrder.UserEmail,
		From:      newOrder.From,
		To:        newOrder.To,
	}

	ctx := context.Background()
	resp, err := handlers.NewCreateOrderHandler(api.ordersRepo, api.roomAvailabilityRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newOrder)
	if err != nil {
		log.Errorf("createOrder - %s", err.Error())
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	log.Info("Order successfully created:\n\tOrder: %v\n\tID: %d", newOrder, resp.OrderID)
}

func (api *BookingApi) SetPreparedData() {
	api.roomAvailabilityRepo.SetPreparedData()
}
