package employee

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	Name   string
	Age    int
	Gender *bool
}
