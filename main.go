package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	employee "github.com/saket3199/GORM-Training/Employee"
)

func main() {
	//connectibf to Database
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	// db.AutoMigrate(&employee.Employee{})

	// employee := employee.Employee{
	// 	Name:   "Xyz",
	// 	Age:    27,
	// 	Gender: false,
	// }
	// // for {
	// db.Debug().Create(&employee)
	// }
	// employees := []employee.Employee{}
	// db.Debug().Table("employees").Find(&employees)
	// fmt.Println(employees)

	// }

	//department related queries
	//migrate database create db from struct auto
	// db.AutoMigrate(&department.Dept{})

	// create department
	// department := department.Dept{
	// 	DeptNo: 10,
	// 	Dep:    "Account",
	// 	LOC:    "Mumbai",
	// }

	// db.Create(&department)

	//fetch whole department
	// departments := []department.Dept{}
	// db.Debug().Table("depts").Find(&departments)
	// fmt.Println(departments)

	// db.Where("Dep = ?", "Accounting").Find(&departments)
	// fmt.Println(departments)
	// db.Where("DEPTNO LIKE ?", "10").Delete(&department.Dept{})
	// db.Unscoped().Delete(&department.Dept{})
	// db.Debug().Where("Name = ?", "Xyz").Delete(&employee.Employee{})
	// db.Debug().Unscoped().Where("name = Xyz").Find(&employee.Employee{})
	// db.Debug().Model(&employee.Employee{}).Updates(map[string]interface{}{"Name": "Saket", "Age": 21, "Gender": true})
	foo := true
	db.Debug().Model(&employee.Employee{}).Updates(&employee.Employee{Name: "Saket", Age: 23, Gender: &foo})
	//
	defer db.Close()
}

// type Dept struct {
// 	DeptNo int
// 	Dep    string
// 	Loc    string
// }
