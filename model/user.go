package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id    int `gorm:"primaryKey"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	IdLine string `json:"idLine"`
	Name string  `json:"Name"`
	Email string `json:"Email"`
	Gender string `json:"Gender"`
	Role string `json:"Role"`
	Isactive bool `json:"IsActive"`
}

