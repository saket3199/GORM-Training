package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/customer_app/Customer"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *gorm.DB) *customer.Customer {
	fmt.Println("Enter the creditientials to login")
start:
	fmt.Println("Enter a Valid Email:")
	fmt.Scan(&email)
	if !valid(email) {
		goto start
	}
	fmt.Println("Enter your pass:")
	fmt.Scan(&pass)
	if pass == "" {
		goto start
	}
	// fmt.Println(email, pass)
	result := []customer.Customer{}
	db.Where("email = ?", email).First(&result)
	// fmt.Println(len(result))
	if len(result) <= 0 {
		fmt.Println("Not A Valid User! Try registering")
		Register(db)

	}
	if err := bcrypt.CompareHashAndPassword([]byte(result[0].UserPass), []byte(pass)); err != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("Invalid Creditintials")

	}

	// fmt.Println(result[0])
	// }

	// var c customer.Customer
	return &result[0]

}
