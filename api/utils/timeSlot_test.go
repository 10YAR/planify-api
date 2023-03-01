package utils_test

import (
	"api/types"
	"api/utils"
	"reflect"
	"testing"
	"time"
)

func TestTimeSlots(t *testing.T) {
	t.Run("Create an array of time slots", func(t *testing.T) {
		// Given
		start := "09:00:00"
		end := "10:00:00"
		duration := 15

		// When
		times := utils.TimeSlots(start, end, duration)

		paris, err := time.LoadLocation("Europe/Paris")
		if err != nil {
			panic(err)
		}

		expectedTimes := []time.Time{
			time.Date(0000, 01, 01, 9, 0, 0, 0, paris),
			time.Date(0000, 01, 01, 9, 15, 0, 0, paris),
			time.Date(0000, 01, 01, 9, 30, 0, 0, paris),
			time.Date(0000, 01, 01, 9, 45, 0, 0, paris),
		}

		// Then
		if !reflect.DeepEqual(expectedTimes, times) {
			t.Errorf("Expected timeSlots to be %v, but got %v", expectedTimes, times)
		}
	})
}

func TestGetTimeSlotsOfAShop(t *testing.T) {
	t.Run("Get time slots from availabilities of shop ", func(t *testing.T) {
		// Given
		availabilities := []types.ShopAvailability{
			{
				DayOfWeek: "saturday",
				Duration:  30,
				StartTime: "09:00:00",
				EndTime:   "10:00:00",
			},
		}

		paris, err := time.LoadLocation("Europe/Paris")
		if err != nil {
			panic(err)
		}

		expectedAvailabilities := []types.ShopAvailabilityWithTimeSlots{
			{
				DayOfWeek: "saturday",
				Duration:  30,
				StartTime: "09:00:00",
				EndTime:   "10:00:00",
				TimeSlots: []time.Time{
					time.Date(0000, 01, 01, 9, 0, 0, 0, paris),
					time.Date(0000, 01, 01, 9, 30, 0, 0, paris),
				},
			},
		}

		// When
		availabilitiesWithTimeSlots := utils.GenerateTimeSlotsOfAShop(availabilities)

		// Then
		if !reflect.DeepEqual(expectedAvailabilities, availabilitiesWithTimeSlots) {
			t.Errorf("Expected availabilities to be %v, but got %v", expectedAvailabilities, availabilitiesWithTimeSlots)
		}
	})
}
