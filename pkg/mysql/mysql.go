package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connection to database
func DatabaseInit() {
	var err error
	dsn := "root:dtlWGDc0unEs0Gve90bS@tcp(containers-us-west-125.railway.app:6140)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to Database")
}
