package model

import (
	"time"

)

type Schedule struct {

	ID            string `gorm:"primaryKey"`
	ScheduleStart time.Time
	ScheduleEnd   time.Time

	EmployeeID string `gorm:"column:employee_id;->;<-:crete"`

	BusID string `gorm:"column:bus_id;->;<-:crete"`

	// BusRefer      string `gorm:"column:Bus_refer"`
	Employee Employee `gorm:"->"`
	Bus Bus  `gorm:"->"`
}


