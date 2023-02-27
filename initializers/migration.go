package initializers

import (
	"server/model"
	"time"
)

func Migration() {
	// CreateDB()
	// DB.Migrator().CreateTable(&model.Bus{}, &model.Employee{}, &model.Schedule{})
	DB.Migrator().CreateTable(&model.Schedule{})
	// CreateBusScheDB()
	// CreateEmpScheDB()
	CreateScheduleDB()
	// DB.Migrator().CreateTable(&model.Employee{}, &model.Bus{}, &model.Schedule{})

	if DB != nil {

		DB.AutoMigrate(&model.User{})

	}

}

func CreateEmpScheDB() {
	DB.Association("employees")
	DB.Create(&model.Employee{
		ID:             "EmployeeID",
		FirstName:      "EmployeeFirstName",
		LastName:       "EmployeeLastName",
		Gender:         "EmployeeGender",
		Age:            0,
		ContactAddress: "EmployeeContactAddress",
		EmployeeEmail:  "EmployeeEmail",
	})
}

func CreateBusScheDB() {
	DB.Association("buses")
	DB.Create(&model.Bus{
		ID:        "BusID",
		BusNumber: "BusNumber",
		BusStop:   "BusStop",
		BusStart:  "BusStart",
	})
}

func CreateScheduleDB() {
	DB.Association("employees")
	DB.Association("schedules")
	DB.Association("buses")
	DB.Create(&model.Schedule{
		ID:            "ScheduleID",
		ScheduleStart: time.Time{},
		ScheduleEnd:   time.Time{},
	})
}

// func CreateDB() {

// 	DB.Association("buses")
// 	DB.Association("employees")
// 	DB.Create(&model.Employee{
// 		ID:             "EmpID",
// 		FirstName:      "EmpFirstName",
// 		LastName:       "EmpLastName",
// 		Gender:         "EmpGender",
// 		Age:            0,
// 		ContactAddress: "EmpContactAddress",
// 		EmployeeEmail:  "EmpEmployeeEmail",
// 		Schedule: []model.Schedule{
// 			{
// 				ID:            "ScheduleID",
// 				ScheduleStart: time.Time{},
// 				ScheduleEnd:   time.Time{},
// 				EmployeeRefer: "EmpRefer",
// 				Bus: []model.Bus{
// 					{
// 						ID:            "BusID",
// 						BusStop:       "BusStop",
// 						BusStart:      "BusStart",
// 						BusEnd:        "BusEnd",
// 						ScheduleRefer: "ScheduleRefer",
// 					},
// 				},
// 			},
// 		},
// 	})
// }
