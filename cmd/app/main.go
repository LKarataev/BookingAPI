package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/LKarataev/BookingAPI/internal/logger"
	"github.com/LKarataev/BookingAPI/internal/service"
)

func main() {
	api := service.NewInMemoryBookingApi()
	api.SetPreparedData()

	mux := api.ConfigureRouter()

	log := logger.GetLogger()
	log.Info("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Info("Server closed")
	} else if err != nil {
		log.Errorf("Server failed: %s", err)
		os.Exit(1)
	}
}
