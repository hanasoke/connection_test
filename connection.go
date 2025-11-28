package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func run() {
	config := DatabaseConfig{
		Username: "root",
		Password: "", // Default XAMPP MySQL tanpa password
		Host:     "127.0.0.1",
		Port:     "3306",
		Database: "testdb",
	}

	// Format connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Ggal membuka koneksi:", err)
	}
	defer db.Close()
}
