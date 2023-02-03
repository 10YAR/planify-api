package types

type Appointment struct {
	ID              int    `json:"id"`
	CustomerName    string `json:"customer_name"`
	AppointmentDate string `json:"appointment_date"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	Status          int    `json:"status"`
	ShopId          int    `json:"shop_id"`
}

type AppointmentInterface interface {
	GetAppointments() ([]Appointment, error)
	GetAppointment(id int) (Appointment, error)
	CreateAppointment(appointment Appointment) (Appointment, error)
	UpdateAppointment(id int, appointment Appointment) (Appointment, error)
	DeleteAppointment(id int) error
}
