// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/LKarataev/BookingAPI/internal/logger"
	"github.com/LKarataev/BookingAPI/internal/service"
)

func main() {
	logger := logger.NewApiLogger()
	api := service.NewBookingApi(logger)
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
