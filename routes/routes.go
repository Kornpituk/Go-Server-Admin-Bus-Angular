package routes

import (
	"server/controllers"

	// "github.com/labstack/echo/"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	e *echo.Echo
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.CORS())

	// User Router
	e.GET("/user", controllers.GetAllUsers)
	e.POST("/user", controllers.CreatedUser)
	e.GET("/user/name/:id", controllers.GetUserByName)
	e.GET("/user/:id", controllers.GetUserByID)
	e.PUT("/user/:id", controllers.EditedUser)
	e.DELETE("/user/:id", controllers.DeletedUser)

	//Role Router
	e.GET("/role", controllers.GetAllRoles)
	e.POST("/role", controllers.CreatedRole)


	//Schedule Router
	e.GET("/schedule",controllers.GetAllSchedule)
	e.GET("/schedule/emp",controllers.GetAllScheduleByEmp)

	//Employee Router
	e.GET("/employee", controllers.GetAllEmployee)

	//Bus Router
	e.GET("/bus", controllers.GetAllBusByEmp)

	return e
}

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	e := echo.New()

// 	e.Use(middleware.CORS())

// 	e.GET("/user", controllers.GetAllUsers)
// 	e.POST("/user", controllers.CreatedUser)
// 	e.GET("/user/:id", controllers.GetUser)
// 	e.PUT("/user/:id", controllers.EditedUser)
// 	e.DELETE("/user/:id", controllers.DeletedUser)

// 	e.GET("/role", controllers.GetAllRoles)
// 	e.POST("/user", controllers.CreatedRole)

// 	e.ServeHTTP(w, r)
// }
