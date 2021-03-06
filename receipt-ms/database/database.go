package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

func SetupDatabase() {

	var err error
	// DbConn, err = sql.Open("mysql", "root:Password1@tcp(db-receipt:3306)/receiptdb")
	DbConn, err = sql.Open("mysql", "root:M@yz13mysJ3r3$$@tcp(127.0.0.1:3306)/receiptdb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
