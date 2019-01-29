package models

import (
	"github.com/jinzhu/gorm"
)

type Manifacturer struct {
	gorm.Model
	Name string `db:"name" json:"name"`
	Cars []Car  `gorm:"foreignkey:ManifacturerID"`
}

func AllBrands() []Manifacturer {
	var manifacturer []Manifacturer
	db.Find(&manifacturer)

	return manifacturer
}

func AllCarsByManifacturer(name string) []Car {

	var manifacturer Manifacturer
	db.Where("name = ?", name).Find(&manifacturer)

	return manifacturer.Cars
}
