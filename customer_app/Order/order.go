package order

import (
	"github.com/saket3199/GORM-Training/customer_app/Model"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Model.Model
	CustomerID  uuid.UUID `gorm:"ForeignKey:CustomerID" type:"uuid"`
	ItemName    string
	ItemDesc    string
	Quantity    int
	CostPerUnit float64
	IsPaid      *bool
}

func New(ID uuid.UUID, itemName, itemDesc string, quantity int, costPerUnit float64, ispaid *bool) *Order {
	return &Order{
		CustomerID:  ID,
		ItemName:    itemName,
		ItemDesc:    itemDesc,
		Quantity:    quantity,
		CostPerUnit: costPerUnit,
		IsPaid:      ispaid,
	}
}
