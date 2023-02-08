package types

type Shop struct {
	ID        int
	ShopName  string
	Address   string
	CreatedAt string
	UserId    int
}

type ShopInfos struct {
	ShopName string
	Address  string
}

type ShopInfosAvailabilitiesAndAppointments struct {
	ShopName       string
	Address        string
	Availabilities []ShopAvailabilityWithTimeSlots
	Appointments   []AppointmentDateTimeInfos
}
