package controllers

import (
	"net/http"
	"server/initializers"
	"server/model"

	"github.com/labstack/echo"
)

// Retrieve user list with eager loading credit cards
func GetAllBusByEmp(c echo.Context) error {
	var buses []model.Bus
	res := initializers.DB.Model(&buses).Preload("Employee").Preload("Schedule").Find(&buses)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK, buses)
}

// func GetAllEmployee(c echo.Context) error {
// 	var employees []model.Employee
// 	res := initializers.DB.Model(&employees).Preload("Schedule").Find(&employees)

// 	if res.Error != nil {
// 		return c.JSON(http.StatusInternalServerError, res.Error)
// 	}

// 	return c.JSON(http.StatusOK, employees)
// }
