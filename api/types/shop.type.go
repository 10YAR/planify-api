package types

type Shop struct {
	ID int
	ShopInfos
	CreatedAt string
	UserId    int
}

type ShopInfos struct {
	ShopName    string
	Description string
	Address     string
	PhoneNumber string
}

type ShopAvailabilities struct {
	Shop
	Availabilities []ShopAvailabilityWithShopId
}

type ShopInfosAvailabilitiesAndAppointments struct {
	ShopInfos
	Availabilities []ShopAvailabilityWithTimeSlots
	Appointments   []Appointment
}
