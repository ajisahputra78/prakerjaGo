package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() (*sql.DB, error) {
	// Konfigurasi koneksi MySQL
	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbPort := "3308"
	dbName := "uk-golang"

	// Buat string koneksi
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Buat koneksi ke MySQL
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// Tes koneksi ke MySQL
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
