package routes

import (
	"github.com/labstack/echo"
	"github.com/username/myapp/app/view"
)

func InitHomeRoutes(e *echo.Echo) {
	e.GET("/", view.Home)
}
