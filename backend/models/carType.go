package models

import (
	"github.com/jinzhu/gorm"
)

type CarType struct {
	gorm.Model
	Name string `db:"name" json:"name"`
	Cars []Car  `gorm:"foreignkey:CarTypeID"`
}

func AllCarsByType(name string) []Car {

	var manifacturer Manifacturer
	db.Where("name = ?", name).Find(&manifacturer)

	return manifacturer.Cars
}
