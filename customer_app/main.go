package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/customer_app/Customer"
	order "github.com/saket3199/GORM-Training/customer_app/Order"
	service "github.com/saket3199/GORM-Training/customer_app/Service"
)

func main() {

	//connectibf to Database
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&customer.Customer{})

	db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
login:
	fmt.Println("Welcome Customer to the Shop...")
	fmt.Println("Please Login or Register to Continue")
	var loginType string
	var cust *customer.Customer
	fmt.Println("1 to Login \n2 to Register")
	reader := bufio.NewReader(os.Stdin)
	loginType, _ = reader.ReadString('\n')
	number, _ := strconv.ParseUint(strings.TrimSpace(loginType), 10, 32)
	if number == 1 {
		cust = service.Login(db)
	} else if number == 2 {
		service.Register(db)
	} else {
		goto login
	}
	fmt.Println(cust)
	// Compare the stored hashed password, with the hashed version of the password that was received
	// if err = bcrypt.CompareHashAndPassword([]byte(c.UserPass), []byte(c.UserPass)); err != nil {
	// If the two passwords don't match, return a 401 status
	// w.WriteHeader(http.StatusUnauthorized)
	// }
	processOrder(db, cust)
	// }

}

func processOrder(db *gorm.DB, c *customer.Customer) {
	reader := bufio.NewReader(os.Stdin)
	for {
		var status string
		fmt.Println("Dear ", c.Fname, c.Lname, ": \n Enter 1 for adding item to cart \n Enter 2 for Order history"+
			"\n Enter 3 for cart \n Enter 4 for checkout"+
			"\n Enter 5 for exit")
		// fmt.Scan(&status)
		status, _ = reader.ReadString('\n')
		number, _ := strconv.ParseUint(strings.TrimSpace(status), 10, 32)
		switch number {
		case 1:
			service.AddItemToCart(db, c)
		case 2:
			service.OrderHistory(db, c)
		case 3:
			service.Usercart(db, c)
		case 4:
			service.INVoice(*c, db)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("default")
		}
	}
}
