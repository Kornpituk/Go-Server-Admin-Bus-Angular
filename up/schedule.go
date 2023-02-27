package model

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	ScheduleStart time.Time
	ScheduleEnd   time.Time

	EmployeeRefer string `gorm:"column:employee_refer"`
	// BusRefer      string `gorm:"column:Bus_refer"`


	Bus []Bus `gorm:"foreignKey:BusRefer"`

}
