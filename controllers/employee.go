package controllers

import (
	"net/http"
	"server/initializers"
	"server/model"

	"github.com/labstack/echo"
)

func GetAllEmployee(c echo.Context) error {
	var employees []model.Employee
	res := initializers.DB.Model(&employees).Preload("Schedule").Find(&employees)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK, employees)
}
