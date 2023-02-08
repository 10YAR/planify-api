package types

type ShopAvailability struct {
	DayOfWeek string
	Duration  int
	StartTime string
	EndTime   string
}

type ShopAvailabilityWithTimeSlots struct {
	DayOfWeek string
	Duration  int
	StartTime string
	EndTime   string
	TimeSlots []TimeSlot
}
