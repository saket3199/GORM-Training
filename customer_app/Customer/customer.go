package customer

import (
	"github.com/saket3199/GORM-Training/customer_app/Model"
)

type Customer struct {
	Model.Model
	Email    string `gorm:"unique;not null"`
	UserPass string
	Fname    string
	Lname    string
	Age      int
	IsMale   *bool
}

func New(email, userPass, fname, lname string, age int, isMale *bool) *Customer {
	return &Customer{
		Email:    email,
		UserPass: userPass,
		Fname:    fname,
		Lname:    lname,
		Age:      age,
		IsMale:   isMale,
	}
}
