package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
)

func GetAllTransactions(db *sql.DB) {
	// Initialize repositories
	transactionRepo := repository.NewTransactionRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	// Fetch all transactions
	transactions, err := transactionRepo.GetAllTransactions()
	if err != nil {
		fmt.Println("Error fetching transactions:", err)
		return
	}

	// Prepare response structure
	type ItemDetail struct {
		ID       uint16         `json:"id"`
		ItemCode string         `json:"item_code"`
		Name     string         `json:"name"`
		Category model.Category `json:"category"`
		Location model.Location `json:"location"`
		Price    int            `json:"price"`
		Stock    int            `json:"stock"`
	}

	type TransactionResponse struct {
		ID           int        `json:"id"`
		Product      ItemDetail `json:"product"`
		Quantity     int        `json:"quantity"`
		MovementType string     `json:"movement_type"`
		Timestamp    string     `json:"timestamp"`
	}

	var result []TransactionResponse

	// Populate the result
	for _, transaction := range transactions {
		product, err := productRepo.FindByID(transaction.ProductID)
		if err != nil {
			fmt.Println("Error fetching product for ID", transaction.ProductID, ":", err)
			continue
		}

		category, err := categoryRepo.FindByID(product.CategoryId)
		if err != nil {
			fmt.Println("Error fetching category for ID", product.CategoryId, ":", err)
			continue
		}

		location, err := locationRepo.FindByID(product.LocationId)
		if err != nil {
			fmt.Println("Error fetching location for ID", product.LocationId, ":", err)
			continue
		}

		itemDetail := ItemDetail{
			ItemCode: product.ItemCode,
			Name:     product.Name,
			Category: *category,
			Location: *location,
			Price:    product.Price,
			Stock:    product.Stock,
		}

		result = append(result, TransactionResponse{
			ID:           transaction.ID,
			Product:      itemDetail,
			Quantity:     transaction.Quantity,
			MovementType: transaction.MovementType,
			Timestamp:    transaction.Timestamp.Format("2006-01-02 15:04:05"),
		})
	}

	response := struct {
		Transactions []TransactionResponse `json:"transactions"`
	}{
		Transactions: result,
	}

	// Convert response to JSON and print it
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling response to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
