package receipt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"receipt-ms/cors"
	"strconv"
	"strings"
)

const receiptsBasePath = "receipts"

func SetupRoutes(apiBasePath string) {

	log.Print("........Receipts Service Routes\n")
	handleReceipt := http.HandlerFunc(receiptHandler)
	handleReceipts := http.HandlerFunc(receiptsHandler)

	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, receiptsBasePath), cors.Middleware(handleReceipt))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, receiptsBasePath), cors.Middleware(handleReceipts))

	log.Print("........Routes Setup Complete")
}

func receiptHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("receipt handler")
	urlPathSegments := strings.Split(r.URL.Path, "receipts/")
	receiptID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodGet:
		receipt, err := getReceipt(receiptID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if receipt == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// return a single receipt
		receiptJSON, err := json.Marshal(receipt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Printf("GET [%s]\n", receiptJSON)
		w.Header().Set("Content-Type", "application/json")
		w.Write(receiptJSON)
	case http.MethodPut:
		fmt.Printf("[Receipt]Put case...\n")
		// update receipt
		var receipt Receipt
		err := json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if receipt.ReceiptID != receiptID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateReceipt(receipt)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		fmt.Printf("[Receipt]Delete case...\n")
		removeReceipt(receiptID)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func receiptsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		receiptList, err := getReceiptList()
		if err != nil {
			fmt.Printf("%s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		receiptsJson, err := json.Marshal(receiptList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("GET [%s]\n", receiptsJson)
		w.Header().Set("Content-Type", "application/json")
		w.Write(receiptsJson)
	case http.MethodPost:
		// add a new product to the list
		var receipt Receipt
		err := json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertReceipt(receipt)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		receiptJSON, err := json.Marshal(receipt)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("POST [%s]\n", receiptJSON)
		w.WriteHeader(http.StatusCreated)

	case http.MethodOptions:
		//relate to CORS -
		return

	}
}
