package types

type User struct {
	ID        int    `validate:"number" json:"id"`
	FirstName string `validate:"required,min=3,max=32" json:"firstName"`
	LastName  string `validate:"required,min=3,max=32" json:"lastName"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required" json:"password"`
	Role      string `validate:"required" json:"role"`
}
