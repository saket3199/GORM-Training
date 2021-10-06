package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/Customer"
	order "github.com/saket3199/GORM-Training/Order"
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

	db.AutoMigrate(&customer.Customer{})
	db.AutoMigrate(&order.Order{})
	db.Model(&order.Order{}).AddForeignKey("CustomerID", "Order(Customer)", "RESTRICT", "RESTRICT")

	var fname, lname, itemName, itemDesc string
	var age, quantity int
	var gender, isPaid bool
	var costPerUnit float64
	var userOrder *order.Order
	fmt.Println("Welcome Customer to the Shop...")
	fmt.Println("Enter your first name:")
	fmt.Scanln(&fname)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lname)
	fmt.Println(fname, " ", lname, " ", "enter your Age:")
	fmt.Scanln(&age)
	fmt.Println(fname, " ", lname, " ", "enter true if you are Male or else false:")
	fmt.Scanln(&gender)
	c := customer.Newc(fname, lname, age, &gender)
	db.Create(c)

	for {
		var status bool
		fmt.Println("Dear ", fname, " ", lname, ": \n Enter 1 for placing order \n Enter 0 for checkout")
		fmt.Scanln(&status)
		if status {
			fmt.Println("Enter the Item name which you want to buy:")
			fmt.Scanln(&itemName)
			fmt.Println("Enter the Description of", itemName, " :")
			fmt.Scan(itemDesc)
			fmt.Println("Enter the number of quantity of", itemName)
			fmt.Scanln(&quantity)
			fmt.Println("Enter cost of per unit of", itemName)
			fmt.Scanln(&costPerUnit)
			fmt.Println("Enter true if you paid", quantity*int(costPerUnit), "in full or else false")
			fmt.Scanln(&isPaid)
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

	rows, err := db.DB().Query("SELECT item_name,item_desc,quantity,cost_per_unit,is_paid FROM orders WHERE customer_id = ?", c.ID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var itemName, itemDesc string
		var quantity int
		var costPerUnit float64
		var isPaid bool
		err = rows.Scan(&itemName, &itemDesc, &quantity, &costPerUnit, &isPaid)
		if err != nil {
			panic(err)
		}
		fmt.Println(itemName, "\t ", itemDesc, "\t", quantity, "\t", costPerUnit, "\t")
		total += float64(quantity) * costPerUnit
	}
	fmt.Println("Total Cost: ", total)
	if err != nil {
		panic(err)
	}
}
