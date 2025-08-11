package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model

	Marca string `gorm:"not null"`
	Modelo string `gorm:"not null"`
	Año int `gorm:"not null"`
	Color string `gorm:"not null"`
	Placa string `gorm:"not null"`
	UserID  uint
}