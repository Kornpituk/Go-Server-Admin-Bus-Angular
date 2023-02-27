package controllers

import (
	"net/http"

	"server/model"

	"server/initializers"

	"github.com/labstack/echo"

	"strconv"
)

var (
	req = 1
)

// function Data Bases Connect SQL
// func init(){
// 	database.ConnectDB()
// }

// Get All Users
func GetAllUsers(c echo.Context) error {
	var user []model.User
	initializers.DB.Find(&user)
	return c.JSON(http.StatusOK, user)
}

// Create User
func CreatedUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	initializers.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

// Get User with ID
func GetUserByName(c echo.Context) error {

	id := c.Param("id")
	var user model.User

	initializers.DB.Where("user_name = ?", id).Find(&user)

	if user.UserName == id {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusOK, "Not Found Name "+c.Param("id")+user.UserName)
}

func GetUserByID(c echo.Context) error {

	id := c.Param("id")
	var user model.User

	initializers.DB.Where("id = ?", id).Find(&user)

	if strconv.Itoa(user.Id) == id {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusBadRequest, "Not Found ID "+c.Param("id")+user.UserName)
}

// Edit User
func EditedUser(c echo.Context) error {

	//Get a hero from the database
	// Get the hero off req body
	id := c.Param("id")
	var User model.User
	c.Bind(&User)
	//Find the hero

	res := initializers.DB.Where("id = ?", id).Model(&User).Updates(model.User{
		Name:     User.Name,
		Email:    User.Email,
		Gender:   User.Gender,
		Role:     User.Role,
		Isactive: User.Isactive,
	})

	//response with the hero
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, res.Error)
	}

	return c.JSON(http.StatusOK, &User)

}

// Delete User
func DeletedUser(c echo.Context) error {

	user := model.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	initializers.DB.Delete(&user, id)

	return c.JSON(http.StatusOK, "Delete Todo Success!"+"  Not Found ID "+c.Param("id"))
}
