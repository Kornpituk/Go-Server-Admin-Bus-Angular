package model

type Employee struct {
	ID     string `gorm:"primaryKey"`
	FirstName      string
	LastName       string
	Gender         string
	Age            int
	ContactAddress string
	EmployeeEmail  string

	// BusRefer string `gorm:"column:bus_refer"`

	Schedule []Schedule `gorm:"foreignKey:EmployeeID"`
}
