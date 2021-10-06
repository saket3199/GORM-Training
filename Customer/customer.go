package customer

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Fname  string
	Lname  string
	Age    int
	Gender *bool
}

func Newc(Fname, Lname string, Age int, Gender *bool) *Customer {
	return &Customer{
		Fname:  Fname,
		Lname:  Lname,
		Age:    Age,
		Gender: Gender,
	}
}
