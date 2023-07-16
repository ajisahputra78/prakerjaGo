package main

import (
	"github.com/labstack/echo"
	"github.com/username/myapp/app/database"
	"github.com/username/myapp/app/model"
	routes "github.com/username/myapp/app/router"
)

func main() {
	// Inisialisasi koneksi database MySQL
	db, err := database.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inisialisasi model buku dengan koneksi database
	model.InitBookModel(db)

	e := echo.New()

	// Routes
	routes.InitHomeRoutes(e)
	routes.InitRoutes(e)

	// Start server
	e.Start(":8000")
}
