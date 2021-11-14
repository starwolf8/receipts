package receipt

import "net/http"

type ReceiptReportFilter struct {
	StoreNameFilter   string `json:"storeName"`
	ProductNameFilter string `json:"productName"`
}

func handleReceiptReport(w http.ResponseWriter, r *http.Request) {}
