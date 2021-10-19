package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/customer_app/Customer"
	order "github.com/saket3199/GORM-Training/customer_app/Order"
)

func Usercart(db *gorm.DB, c *customer.Customer) {
	orders := []order.Order{}
	var total float64 = 0
	db.Where("customer_id = ? AND is_paid = ?", c.ID, false).Find(&orders)
	count := len(orders)
	fmt.Println("Total Items in cart: ", count)
	if count < 1 {
		fmt.Println("Hey", c.Fname, "Your Cart is Empty ")
	} else {
		fmt.Println("Item ID\tItem name\tItem Desc\tQuantity\tCost/unit\t")
		for i, element := range orders {
			fmt.Println(i, "\t ", element.ItemName, "\t ", element.ItemDesc, "\t", element.Quantity, "\t", element.CostPerUnit, "\t")
			total += float64(element.Quantity) * element.CostPerUnit
		}
		var sta bool
		fmt.Println("Total Cost: ", total)

		for {
			fmt.Println("Enter 1 for removing an Item or 0 to continue")
			fmt.Scan(&sta)
			if sta {
			cartID:
				var pro string
				fmt.Println("Enter the Item ID to remove the Item from Cart")
				fmt.Scan(&pro)
				number, err := strconv.ParseUint(pro, 10, 32)
				if err == nil {
					db.Delete(&orders[number])
				} else {
					goto cartID
				}

			} else {
				break
			}
		}
	}
}

func AddItemToCart(db *gorm.DB, c *customer.Customer) {
	fmt.Println("Enter the Item name which you want to buy:")
	fmt.Scan(&itemName)
	fmt.Println("Enter the Description of", itemName, " :")
	fmt.Scan(&itemDesc)
cartItemQ:
	fmt.Println("Enter the number of quantity of", itemName)
	// fmt.Scan(&quantity)
	reader := bufio.NewReader(os.Stdin)
	quants, _ := reader.ReadString('\n')
	number, err := strconv.ParseUint(strings.TrimSpace(quants), 10, 32)
	if err == nil {
		quantity = int(number)
	} else {
		goto cartItemQ
	}
cartItemP:
	fmt.Println("Enter cost of per unit of", itemName)
	// fmt.Scan(&costPerUnit)
	costs, _ := reader.ReadString('\n')
	costItem, err := strconv.ParseUint(strings.TrimSpace(costs), 10, 32)
	if err == nil {
		costPerUnit = float64(costItem)
	} else {
		goto cartItemP
	}
	isPaid = false
	userOrder = order.New(c.ID, itemName, itemDesc, quantity, costPerUnit, &isPaid)
	db.Create(userOrder)

}
