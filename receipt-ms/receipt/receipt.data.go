package receipt

import (
	"context"
	"database/sql"
	"receipt-ms/database"
	"sync"
	"time"
)

var receiptMap = struct {
	sync.RWMutex
	m map[int]Receipt
}{m: make(map[int]Receipt)}

func getReceipt(receiptID int) (*Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	row := database.DbConn.QueryRowContext(ctx, `SELECT receiptID,
	storeName,
	dateOfPurchase,
	totalCost,
	taxRate, 
	totalTax
	FROM receipts
	WHERE receiptID = ? `, receiptID)

	receipt := &Receipt{}
	err := row.Scan(&receipt.ReceiptID,
		&receipt.StoreName,
		&receipt.DateOfPurchase,
		&receipt.TotalCost,
		&receipt.TaxRate,
		&receipt.TotalTax)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return receipt, nil
}
func getReceiptList() ([]Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT receiptID,
	storeName,
	dateOfPurchase,
	totalCost,
	taxRate, 
	totalTax
	FROM receipts`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	receipts := make([]Receipt, 0)
	for results.Next() {
		var receipt Receipt
		results.Scan(&receipt.ReceiptID,
			&receipt.StoreName,
			&receipt.DateOfPurchase,
			&receipt.TotalCost,
			&receipt.TaxRate,
			&receipt.TotalTax)
		receipts = append(receipts, receipt)
	}
	return receipts, nil

}

func updateReceipt(receipt Receipt) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `UPDATE receipts SET
	storeName=?,
	dateofPurchase=?,
	totalCost=CAST(? AS DECIMAL(13,2)),
	taxRate=CAST(? AS DECIMAL(13,2)),
	totalTax=?
	WHERE receiptID=?`,
		receipt.StoreName,
		receipt.DateOfPurchase,
		receipt.TotalCost,
		receipt.TaxRate,
		receipt.TotalTax,
		receipt.ReceiptID)
	if err != nil {
		return err
	}
	return nil
}

func insertReceipt(receipt Receipt) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO receipt
	(storeName, 
		dateOfPurchase, 
		TotalCost,
		TaxRate,
		TotalTax) VALUES (?,?,?,?,?)`,
		receipt.StoreName,
		receipt.DateOfPurchase,
		receipt.TotalCost,
		receipt.TaxRate,
		receipt.TotalTax)
	if err != nil {
		return 0, nil
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil

}

func removeReceipt(receiptID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM receipts where receiptID = ?`, receiptID)
	if err != nil {
		return err
	}
	return nil
}

// func searchForReceiptData(receiptFilter ReceiptReportFilter) ([]Receipt, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	var queryArgs = make([]interface{}, 0)
// 	var queryBuilder strings.Builder
// 	queryBuilder.WriteString(`SELECT
// 		receiptID,
// 		LOWER(storeName),
// 		dateOfPurchase,
// 		totalCost,
// 		taxRate,
// 		totalTax
// 		FROM receipts WHERE `)
// 	if receiptFilter.StoreNameFilter != "" {
// 		queryBuilder.WriteString(`productName LIKE ? `)
// 		queryArgs = append(queryArgs, "%"+strings.ToLower(receiptFilter.StoreNameFilter)+"%")
// 	}
// 	if receiptFilter.ProductNameFilter != "" {
// 		if len(queryArgs) > 0 {
// 			queryBuilder.WriteString(" AND ")
// 		}
// 		queryBuilder.WriteString(`productName LIKE ? `)
// 		queryArgs = append(queryArgs, "%"+strings.ToLower(receiptFilter.ProductNameFilter)+"%")
// 	}

// 	results, err := database.DbConn.QueryContext(ctx, queryBuilder.String(), queryArgs...)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, err
// 	}
// 	defer results.Close()
// 	receipts := make([]Receipt, 0)
// 	for results.Next() {
// 		var receipt Receipt
// 		results.Scan(&receipt.ReceiptID,
// 			&receipt.StoreName,
// 			&receipt.DateOfPurchase,
// 			// &receipt.ProductList,
// 			&receipt.TotalCost,
// 			&receipt.TaxRate,
// 			&receipt.TotalTax)

// 		receipts = append(receipts, receipt)
// 	}
// 	return receipts, nil
// }
