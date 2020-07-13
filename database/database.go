package database

import (
	"database/sql"
	"log"
	"time"
)

var DBConnect *sql.DB

func SetupDatabase() {
	var err error
	DBConnect, err = sql.Open("mysql", "aleksa:qweQWE1_@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DBConnect.SetMaxOpenConns(4)
	DBConnect.SetMaxIdleConns(4)
	DBConnect.SetConnMaxLifetime(60 * time.Second)
}
