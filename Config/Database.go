package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/vaibhav/practice/projectpharmamanagementsystem/Models"
)

var DB *gorm.DB
var err error

const SC = "user.db"

func DbSetup() {
	DB, err = gorm.Open("sqlite3", SC)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed connect of db")
	}

	DB.AutoMigrate(&Models.Manager{})
	DB.AutoMigrate(&Models.Pharmacist{})
	DB.AutoMigrate(&Models.Salesman{})
	DB.AutoMigrate(&Models.Medices{})
	
}
