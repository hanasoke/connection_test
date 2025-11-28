package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func maind() {
// 	// Konfigurasi koneksi database
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Tesk koneksi
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Berhasil terhubung ke database!")
// }
