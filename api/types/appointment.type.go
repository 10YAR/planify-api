package types

type Appointment struct {
	ID              int    `validate:"number" json:"id"`
	CustomerName    string `validate:"required,min=3,max=32" json:"customer_name"`
	AppointmentDate string `validate:"required,len=10" json:"appointment_date"`
	StartTime       string `validate:"required,len=19" json:"start_time"`
	EndTime         string `validate:"required,len=19" json:"end_time"`
	Status          int    `validate:"required,number" json:"status"`
	ShopId          int    `validate:"required,number" json:"shop_id"`
}
