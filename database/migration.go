package database

import (
	"dewe/models"
	"dewe/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Country{},
		&models.Trip{},
		&models.Transaction{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
