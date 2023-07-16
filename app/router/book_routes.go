package routes

import (
	"github.com/labstack/echo"
	"github.com/username/myapp/app/controller"
)

func InitRoutes(e *echo.Echo) {
	// Book routes
	e.GET("/api/books", controller.GetAllBooks)
	e.GET("/api/books/:id", controller.GetBookByID)
	e.POST("/api/books/create", controller.CreateBook)
	e.PUT("/api/books/:id/update", controller.UpdateBook)
	e.DELETE("/api/books/:id/delete", controller.DeleteBook)
}
