package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	Code string `gorm:"primaryKey"`
	Name string `json:"Name"`
}
