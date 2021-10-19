package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/customer_app/Customer"
	order "github.com/saket3199/GORM-Training/customer_app/Order"
)

func OrderHistory(db *gorm.DB, c *customer.Customer) {
	orders := []order.Order{}
	db.Where("customer_id = ? AND is_paid = ?", c.ID, true).Find(&orders)
	count := len(orders)
	fmt.Println("Total Items Placed: ", count)
	fmt.Println("Item name\tItem Desc\tQuantity\tCost/unit\t")
	for _, element := range orders {
		fmt.Println(element.ItemName, "\t ", element.ItemDesc, "\t", element.Quantity, "\t", element.CostPerUnit, "\t")
	}

}
