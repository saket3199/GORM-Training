package department

import "github.com/jinzhu/gorm"

type Dept struct {
	gorm.Model
	DeptNo int
	Dep    string
	LOC    string
}
