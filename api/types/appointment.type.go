package types

type Appointment struct {
	ID           int    `validate:"number" json:"id"`
	CustomerName string `validate:"required,min=3,max=32" json:"customer_name"`
	AppointmentDateTimeInfos
	ShopId int    `validate:"required,number" json:"shop_id"`
	UserId int    `validate:"" json:"user_id"`
	Email  string `validate:"required,email" json:"user_email"`
}

type AppointmentDateTimeInfos struct {
	AppointmentDate     string `validate:"required,len=10" json:"appointment_date"`
	AppointmentTime     string `validate:"required,min=7,max=8" json:"appointment_time"`
	AppointmentDateTime string `validate:"" json:"appointment_date_time"`
}

type AppointmentWithShop struct {
	Appointment
	ShopName string `validate:"" json:"shop_name"`
}
