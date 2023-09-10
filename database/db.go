package database

import (
	
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var PersonDB *gorm.DB

func LoadDB() {
	var err error
	dbUri := "student:student@tcp(127.0.0.1:3306)/person-service?charset=utf8mb4&parseTime=True&loc=Local"
	PersonDB, err = gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to db")
	}

}

