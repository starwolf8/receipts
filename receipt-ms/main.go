package main

import (
	"log"
	"receipt-ms/database"
	"receipt-ms/receipt"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {

	log.Print("Initializing Receipts Microservice\n")
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5001", nil)
}
