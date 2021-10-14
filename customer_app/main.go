package main

import (
	"fmt"
	"log"
	"net/mail"

	_ "github.com/go-sql-driver/mysql"
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

func main() {

	//connectibf to Database
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	defer db.Close()

	db.AutoMigrate(&customer.Customer{})
	db.AutoMigrate(&order.Order{})
	db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "RESTRICT")
login:
	fmt.Println("Welcome Customer to the Shop...")
	fmt.Println("Please Login or Register to Continue")
	var loginType uint
	fmt.Println("1 to Login \n2 to Register")
	fmt.Scan(&loginType)
	if loginType == 1 {
		login(db)
	} else if loginType == 2 {
		register(db)
	} else {
		goto login
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	// if err = bcrypt.CompareHashAndPassword([]byte(c.UserPass), []byte(c.UserPass)); err != nil {
	// If the two passwords don't match, return a 401 status
	// w.WriteHeader(http.StatusUnauthorized)
	// }

}
func login(db *gorm.DB) {
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
	fmt.Println(len(result))
	if len(result) <= 0 {
		fmt.Println("Not A Valid User! Try registering")
		register(db)

	}
	if err := bcrypt.CompareHashAndPassword([]byte(result[0].UserPass), []byte(pass)); err != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("Invalid Creditintials")
		return
	}

	// fmt.Println(result[0])
	// }

	// var c customer.Customer

	processOrder(db, &result[0])

}
func register(db *gorm.DB) {
	fmt.Println("Welcome Customer Register Yourself")
start:
	fmt.Println("Enter a valid Email Email:")
	fmt.Scan(&email)
	if !valid(email) {
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
	c := customer.Newc(email, pass, fname, lname, age, &gender)
	db.Create(c)
	// processOrder(db, c)
	login(db)
}
func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func processOrder(db *gorm.DB, c *customer.Customer) {
	for {
		var status bool
		fmt.Println("Dear ", c.Fname, " ", c.Lname, ": \n Enter 1 for placing order \nEnter 0 for checkout")
		fmt.Scan(&status)
		if status {
			fmt.Println("Enter the Item name which you want to buy:")
			fmt.Scan(&itemName)
			fmt.Println("Enter the Description of", itemName, " :")
			fmt.Scan(itemDesc)
			fmt.Println("Enter the number of quantity of", itemName)
			fmt.Scan(&quantity)
			fmt.Println("Enter cost of per unit of", itemName)
			fmt.Scan(&costPerUnit)
			// fmt.Println("Enter true if you paid", quantity*int(costPerUnit), "in full or else false")
			// fmt.Scan(&isPaid)
			isPaid = false
			userOrder = order.New(c.ID, itemName, itemDesc, quantity, costPerUnit, &isPaid)
			db.Create(userOrder)
		} else {
			inVoice(*c, db)
			break
		}
	}
}

func inVoice(c customer.Customer, db *gorm.DB) {
	var total float64 = 0
	fmt.Println("Name : ", c.Fname, " ", c.Lname)
	fmt.Println("Age  : ", c.Age)
	fmt.Println("IsMale : ", *c.Gender)
	fmt.Println("Item name\tItem Desc\tQuantity\tCost/unit\t")

	orders := []order.Order{}
	db.Where("customer_id = ? AND is_paid = ?", c.ID, false).Find(&orders)

	for _, element := range orders {
	
		fmt.Println(element.ItemName, "\t ", element.ItemDesc, "\t", element.ItemDesc, "\t", element.CostPerUnit, "\t")
		total += float64(element.Quantity) * element.CostPerUnit
		db.Model(&element).Update("is_paid", true)
	}
	fmt.Println("Total Cost: ", total)
	fmt.Println("Confirming your Order Thank you for shopping with us")
	
}
