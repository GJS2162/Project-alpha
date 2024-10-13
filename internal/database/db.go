package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() error {
	// Use the correct connection string for the MySQL database
	dsn := "sql12737462:ClNqFAxSmT@tcp(sql12.freesqldatabase.com:3306)/sql12737462"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error initializing MySQL client: %v", err)
	}

	// Check if the database is reachable
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to MySQL: %v", err)
	}

	fmt.Println("Successfully connected to MySQL")
	return nil
}
