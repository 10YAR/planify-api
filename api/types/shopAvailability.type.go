package types

type ShopAvailability struct {
	DayOfWeek string
	TimeRange int
	StartTime string
	EndTime   string
}

type ShopInfosWithAvailabilities struct {
	ShopName       string
	Address        string
	Availabilities []ShopAvailability
}
