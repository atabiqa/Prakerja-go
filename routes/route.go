package routes

import (
	usercontroller "prakerja11/controllers/user"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/users", usercontroller.GetUserController)
}
