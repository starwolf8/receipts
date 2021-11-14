package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

func SetupDatabase() {

	log.Print("........Database\n")
	var err error
	DbConn, err = sql.Open("mysql", "<user>:<pw>@tcp(127.0.0.1:3306)/receiptdb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
