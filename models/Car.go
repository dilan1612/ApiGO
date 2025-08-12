package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model

	Marca  string `gorm:"not null" json:"marca"`
	Modelo string `gorm:"not null" json:"modelo"`
	Año    int    `gorm:"not null" json:"year"`
	Color  string `gorm:"not null" json:"color"`
	Placa  string `gorm:"not null;unique" json:"placa"`
	UserID uint   `json:"user_id"`
}
