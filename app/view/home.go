package view

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Uji Keterampilan Golang Pijar Betch 6")
}