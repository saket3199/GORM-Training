package customer

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	UserPass string
	Fname    string
	Lname    string
	Age      int
	Gender   *bool
}

func Newc(Email, UserPass, Fname, Lname string, Age int, Gender *bool) *Customer {
	return &Customer{
		Email:    Email,
		UserPass: UserPass,
		Fname:    Fname,
		Lname:    Lname,
		Age:      Age,
		Gender:   Gender,
	}
}
