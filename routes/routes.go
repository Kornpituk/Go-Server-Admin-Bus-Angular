package routes

import (
	"net/http"
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

	e.GET("/user", controllers.GetAllUsers)
	e.POST("/user", controllers.CreatedUser)
	e.GET("/user/:id", controllers.GetUser)
	e.PUT("/user/:id", controllers.EditedUser)
	e.DELETE("/user/:id", controllers.DeletedUser)


	return e
}


func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/user", controllers.GetAllUsers)
	e.POST("/user", controllers.CreatedUser)
	e.GET("/user/:id", controllers.GetUser)
	e.PUT("/user/:id", controllers.EditedUser)
	e.DELETE("/user/:id", controllers.DeletedUser)

	e.ServeHTTP(w, r)
}
