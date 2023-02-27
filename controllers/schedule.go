package controllers

import (
	"net/http"
	"server/initializers"
	"server/model"

	"github.com/labstack/echo"
)

// Retrieve user list with eager loading credit cards
func GetAllScheduleByEmp(c echo.Context) error {
	var employees []model.Employee
	err := initializers.DB.Model(&employees).Preload("Schedule").Find(&employees).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, err)
}

func GetAllSchedule(c echo.Context) error {
	var schedules []model.Schedule
	res := initializers.DB.Model(&schedules).Preload("Employee").Preload("Bus").Find(&schedules)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK,schedules)
}


// func CreateSchedule(c echo.Context) error {
// 	var emp model.Employee
// 	var schedule model.Schedule

// 	if err := c.Bind(schedule); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	if err := c.Bind(emp); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	res := initializers.DB.Create(&model.Employee{
// 		ID:             emp.ID,
// 		FirstName:      emp.FirstName,
// 		LastName:       emp.LastName,
// 		Gender:         emp.Gender,
// 		Age:            emp.Age,
// 		ContactAddress: emp.ContactAddress,
// 		EmployeeEmail:  emp.EmployeeEmail,
// 		BusRefer:       emp.BusRefer,
// 		Schedule:       []model.Schedule{
// 			{
// 				ID:    schedule.ID,
// 				ScheduleStart: schedule.ScheduleStart,
// 				ScheduleEnd: schedule.ScheduleEnd,
// 				EmployeeRefer: schedule.EmployeeRefer,
// 			},
// 		},
// 	})
// 	return c.JSON(http.StatusCreated, res)
// }

