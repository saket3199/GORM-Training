package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/customer_app/Customer"
	order "github.com/saket3199/GORM-Training/customer_app/Order"
)

func INVoice(c customer.Customer, db *gorm.DB) {
	orders := []order.Order{}
	db.Where("customer_id = ? AND is_paid = ?", c.ID, false).Find(&orders)
	if len(orders) < 1 {
		fmt.Println("Hey", c.Fname, "Your Cart is Empty ")
		// processOrder(db, &c)
	} else {
		var total float64 = 0
		fmt.Println("Name : ", c.Fname, " ", c.Lname)
		fmt.Println("Age  : ", c.Age)
		fmt.Println("IsMale : ", *c.IsMale)

		fmt.Println("Item name\tItem Desc\tQuantity\tCost/unit\t")
		for _, element := range orders {

			fmt.Println(element.ItemName, "\t ", element.ItemDesc, "\t", element.Quantity, "\t", element.CostPerUnit, "\t")
			total += float64(element.Quantity) * element.CostPerUnit
			db.Model(&element).Update("is_paid", true)
		}
		fmt.Println("Total Cost: ", total)
		fmt.Println("Confirming your Order Thank you for shopping with us")

	}
}
