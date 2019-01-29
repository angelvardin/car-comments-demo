package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	//Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

//InitDB initialize DB
func InitDB(connectionString string) {
	var err error
	fmt.Println("Go MySQL Tutorial")
	initialMigration(connectionString)
	db, err = gorm.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	if db == nil {
		panic("Null db")
	} else {
		fmt.Println("Not null db")
	}
}
func initialMigration(connectionString string) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Car{})
	db.AutoMigrate(&CarType{})
	db.AutoMigrate(&Manifacturer{})
}

func closeDb() {
	defer db.Close()
}
