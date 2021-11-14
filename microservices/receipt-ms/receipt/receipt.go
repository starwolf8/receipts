package receipt

type Receipt struct {
	ReceiptID      int    `json:"receiptID"`
	StoreName      string `json:"storeName"`
	DateOfPurchase string `json:"dateOfPurchase"`
	// ProductList    []Product `json:"productList"`
	TotalCost string `json:"totalCost"`
	TaxRate   string `json:"taxRate"`
	TotalTax  string `json:"totalTax"`
}

type Product struct {
	ReceiptID int    `json:"receiptID"`
	ProductID int    `json:"productID"`
	Name      string `json:"productName"`
	Cost      string `json:"productCost"`
	Quantity  string `json:"productQuantity"`
}
