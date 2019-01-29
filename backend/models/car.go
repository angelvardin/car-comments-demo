package models

import (
	"github.com/jinzhu/gorm"
)

type Car struct {
	gorm.Model
	CarModel       string `db:"model" json:"model"`
	Engine         string `db:"engine" json:"engine"`
	Power          uint16 `db:"power" json:"power"`
	CarType        CarType
	CarTypeID      uint
	Manifacturer   Manifacturer
	ManifacturerID uint
	FuelType       string `db:"fuel_type" json:"fuel_type"`
}

func GetAllCars() []Car {

	var cars []Car
	db.Find(&cars)
	for i, _ := range cars {
		db.Model(cars[i]).Related(&cars[i].CarType)
		db.Model(cars[i]).Related(&cars[i].Manifacturer)
	}

	return cars

}

func Contains(name string) []Car {

	var cars []Car
	db.Where("name LIKE ?", "%"+name+"%").Find(&cars)
	return cars
}
