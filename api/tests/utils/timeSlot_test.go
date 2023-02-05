package tests

import (
	"api/types"
	"api/utils"
	"reflect"
	"testing"
	"time"
)

func TestTimeSlot(t *testing.T) {
	t.Run("Create a time slot", func(t *testing.T) {
		// Given
		start := time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)
		duration := 15 * time.Minute

		// When
		timeSlot := utils.TimeSlot(start, duration)

		expectedStart := start
		expectedEnd := start.Add(duration)

		// Then
		if timeSlot.Start != expectedStart {
			t.Errorf("Expected start time to be %v, but got %v", expectedStart, timeSlot.Start)
		}
		if timeSlot.End != expectedEnd {
			t.Errorf("Expected end time to be %v, but got %v", expectedEnd, timeSlot.End)
		}
	})
}

func TestTimeSlots(t *testing.T) {
	t.Run("Create an array of time slots", func(t *testing.T) {
		// Given
		start := "09:00:00"
		end := "17:00:00"
		duration := 30

		layout := "09:00:00"
		startTime, _ := time.Parse(layout, start)
		convertedDuration := time.Duration(duration) * time.Minute

		// When
		times := utils.TimeSlots(start, end, duration)

		expectedStart := startTime
		expectedEnd := startTime.Add(convertedDuration)

		// Then
		for i, timeSlots := range times {
			if timeSlots.Start != expectedStart {
				t.Errorf("Expected start time for slot %d to be %v, but got %v", i, expectedStart, timeSlots.Start)
			}
			if timeSlots.End != expectedEnd {
				t.Errorf("Expected end time for slot %d to be %v, but got %v", i, expectedEnd, timeSlots.End)
			}

			expectedStart = expectedEnd
			expectedEnd = expectedEnd.Add(convertedDuration)
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

		expectedAvailabilities := []types.ShopAvailabilityWithTimeSlots{
			{
				DayOfWeek: "saturday",
				Duration:  30,
				StartTime: "09:00:00",
				EndTime:   "10:00:00",
				TimeSlots: []types.TimeSlot{
					{
						Start: time.Date(0000, 01, 01, 9, 0, 0, 0, time.UTC), End: time.Date(0000, 01, 01, 9, 30, 0, 0, time.UTC),
					},
					{
						Start: time.Date(0000, 01, 01, 9, 30, 0, 0, time.UTC), End: time.Date(0000, 01, 01, 10, 0, 0, 0, time.UTC),
					},
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
