package order

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	CustomerID  uint `gorm:"ForeignKey:CustomerID"`
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
