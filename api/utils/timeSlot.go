package utils

import (
	"api/types"
	"fmt"
	"time"
)

func TimeSlots(start string, end string, duration int) []time.Time {
	var times []time.Time
	layout := "15:00:00"

	paris, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		panic(err)
	}

	startTime, errStart := time.ParseInLocation(layout, start, paris)
	if errStart != nil {
		fmt.Println("Error parsing startTime:", errStart)
	}

	endTime, errEnd := time.ParseInLocation(layout, end, paris)
	if errEnd != nil {
		fmt.Println("Error parsing endTime:", errEnd)
	}

	durationTime := time.Duration(duration) * time.Minute

	for t := startTime; t.Before(endTime); t = t.Add(durationTime) {
		if t.Add(durationTime).After(endTime) {
			break
		}

		times = append(times, t)
	}
	return times
}

func GenerateTimeSlotsOfAShop(availabilities []types.ShopAvailability) []types.ShopAvailabilityWithTimeSlots {
	var availabilitiesWithTimeSlots []types.ShopAvailabilityWithTimeSlots
	for _, availability := range availabilities {

		availabilityWithTimeSlots := types.ShopAvailabilityWithTimeSlots{
			DayOfWeek: availability.DayOfWeek,
			Duration:  availability.Duration,
			StartTime: availability.StartTime,
			EndTime:   availability.EndTime,
			TimeSlots: TimeSlots(availability.StartTime, availability.EndTime, availability.Duration),
		}

		availabilitiesWithTimeSlots = append(availabilitiesWithTimeSlots, availabilityWithTimeSlots)
	}
	return availabilitiesWithTimeSlots
}
