package model

import (
	"gorm.io/gorm"
) 

type Bus struct {
	gorm.Model

	ID    string `gorm:"primaryKey"`
	BusNumber string 
	BusStop  string
	BusStart string
	BusEnd   string

	ScheduleRefer string `gorm:"column:schedule_refer"`

	// Employee []Employee `gorm:"foreignKey:BusRefer"`
}
