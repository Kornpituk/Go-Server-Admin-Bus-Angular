package controllers

import (
	"net/http"

	"server/model"

	"server/initializers"

	"github.com/labstack/echo"
)

// function Data Bases Connect SQL
// func init(){
// 	database.ConnectDB()
// }

// Get All Roles
func GetAllRoles(c echo.Context) error {
	var role []model.Role
	initializers.DB.Find(&role)


	return c.JSON(http.StatusOK, role)
}

// Create Role
func CreatedRole(c echo.Context) error {
	role := new(model.Role)
	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	initializers.DB.Create(&role)
	return c.JSON(http.StatusCreated, role)
}
