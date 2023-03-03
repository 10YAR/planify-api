package types

import "time"

type ShopAvailability struct {
	DayOfWeek string
	Duration  int
	StartTime string
	EndTime   string
}

type ShopAvailabilityWithShopId struct {
	ShopAvailability
	ShopId int
}

type ShopAvailabilityWithTimeSlots struct {
	DayOfWeek string
	Duration  int
	StartTime string
	EndTime   string
	TimeSlots []time.Time
}
