package handlers

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/LKarataev/BookingAPI/internal/logger"
)

type CreateOrderCase struct {
	Ctx     context.Context
	Request CreateOrderRequest
	Result  int
	Err     error
}

func TestCreateOrder(t *testing.T) {
	cases := []CreateOrderCase{
		CreateOrderCase{
			Ctx: context.Background(),
			Request: CreateOrderRequest{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "guest@mail.ru",
				From:      time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
				To:        time.Date(2024, time.Month(1), 4, 0, 0, 0, 0, time.UTC),
			},
			Result: 4,
			Err:    nil,
		},
		CreateOrderCase{
			Ctx: context.Background(),
			Request: CreateOrderRequest{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "guest@mail.ru",
				From:      time.Date(2024, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
				To:        time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC),
			},
			Result: 4,
			Err:    fmt.Errorf("Data with HotelID=reddison, RoomID=lux, Date=%v not found in application memory", time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.UTC)),
		},
		CreateOrderCase{
			Ctx: context.Background(),
			Request: CreateOrderRequest{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "guest@mail.ru",
				From:      time.Date(2023, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
				To:        time.Date(2023, time.Month(1), 5, 0, 0, 0, 0, time.UTC),
			},
			Result: 4,
			Err:    fmt.Errorf("Data with HotelID=reddison, RoomID=lux, Date=%v not found in application memory", time.Date(2023, time.Month(1), 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	runCreateOrderCases(t, cases)
}

func runCreateOrderCases(t *testing.T, cases []CreateOrderCase) {
	ordersRepo := NewOrdersRepositoryMock()
	roomAvailabilityRepo := NewRoomAvailabilityRepositoryMock()
	logger := logger.NewApiLogger()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expected := item.Result
		expectedErr := item.Err
		result, resultErr := NewCreateOrderHandler(ordersRepo, roomAvailabilityRepo, logger).Handle(item.Ctx, item.Request)

		if expectedErr != nil && resultErr != nil {
			if resultErr.Error() != expectedErr.Error() {
				t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, resultErr, expectedErr)
			}
			continue
		}

		if !(expectedErr == nil && resultErr == nil) {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, resultErr, expectedErr)
			continue
		}

		if result.OrderID != expected {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result.OrderID, expected)
			continue
		}
	}
}
