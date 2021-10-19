package service

import (
	"fmt"
	"net/mail"

	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/customer_app/Customer"
	order "github.com/saket3199/GORM-Training/customer_app/Order"
	"golang.org/x/crypto/bcrypt"
)

var email, pass, fname, lname, itemName, itemDesc string
var age, quantity int
var gender, isPaid bool
var costPerUnit float64
var userOrder *order.Order

func Register(db *gorm.DB) {
	fmt.Println("Welcome Customer Register Yourself")
start:
	fmt.Println("Enter a valid Email Email:")
	fmt.Scan(&email)
	if !valid(email) {
		goto start
	}
	result := []customer.Customer{}
	db.Where("email = ?", email).First(&result)
	if len(result) > 0 {
		fmt.Println("Email Already in Use")
		goto start
	}
	fmt.Println("Enter your pass:")
	fmt.Scan(&pass)
	if pass == "" {
		goto start
	} else {
		hash, err := MakePassword(pass)
		if err != nil {
			return
		}
		pass = hash
	}
	fmt.Println("Enter your first name:")
	fmt.Scan(&fname)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lname)
	fmt.Println(fname, lname, "enter your Age:")
	fmt.Scan(&age)
	fmt.Println(fname, lname, "enter true if you are Male or else false:")
	fmt.Scan(&gender)
	c := customer.New(email, pass, fname, lname, age, &gender)
	db.Create(c)
	// processOrder(db, c)
	Login(db)
}
func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
