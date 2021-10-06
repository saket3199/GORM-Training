package order

import (
	"github.com/jinzhu/gorm"
	customer "github.com/saket3199/GORM-Training/Customer"
)

var ID uint = 1

type Order struct {
	gorm.Model
	Customer    customer.Customer `gorm:"ForeignKey:CustomerID"`
	CustomerID  uint
	ItemName    string
	ItemDesc    string
	Quantity    int
	CostPerUnit float64
	IsPaid      *bool
}

func New(ID uint, ItemName, ItemDesc string, Quantity int, CostPerUnit float64, Ispaid *bool) *Order {
	return &Order{
		CustomerID:  ID,
		ItemName:    ItemName,
		ItemDesc:    ItemDesc,
		Quantity:    Quantity,
		CostPerUnit: CostPerUnit,
		IsPaid:      Ispaid,
	}
}
