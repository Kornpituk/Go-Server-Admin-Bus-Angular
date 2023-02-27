package model

type Bus struct {
	ID        string `gorm:"primaryKey"`
	BusNumber string
	BusStop   string
	BusStart  string

	Schedule []Schedule `gorm:"foreignKey:BusID"`

	// ScheduleRefer string `gorm:"column:schedule_rfer"`

}
