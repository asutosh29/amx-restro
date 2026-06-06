package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/asutosh29/amx-restro/pkg/utils/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() (*sql.DB, error) {
	var err error

	// Get database connection details from environment
	DbHost := config.Db_config.DbHost
	DbUser := config.Db_config.DbUser
	DbPassword := config.Db_config.DbPassword
	Database := config.Db_config.Database
	DbPort := config.Db_config.DbPort

	// Create connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DbUser, DbPassword, DbHost, DbPort, Database)

	// Open database connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configure connection pool settings
	DB.SetMaxOpenConns(25)                 // Maximum number of open connections
	DB.SetMaxIdleConns(5)                  // Maximum number of idle connections
	DB.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Database connected successfully!")
	return DB, nil
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if DB != nil {
		fmt.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}
