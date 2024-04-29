package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/LKarataev/BookingAPI/internal/dao"
	"github.com/LKarataev/BookingAPI/internal/logger"
	"github.com/LKarataev/BookingAPI/internal/service"
)

func main() {
	logger := logger.NewApiLogger()
	ordersRepo := &dao.OrdersRepository{}
	roomAvailabilityRepo := &dao.RoomAvailabilityRepository{}

	api := service.NewBookingApi(ordersRepo, roomAvailabilityRepo, logger)
	api.SetPreparedData()

	mux := api.ConfigureRouter()

	logger.Info("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("Server closed")
	} else if err != nil {
		logger.Errorf("Server failed: %s", err)
		os.Exit(1)
	}
}
