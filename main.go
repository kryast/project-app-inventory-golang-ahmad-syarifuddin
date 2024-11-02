package main

import (
	"fmt"
	"log"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/database"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/handler"
	_ "github.com/lib/pq"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		var endpoint string
		fmt.Print("Masukkan endpoint (create/update/movement/category/location) or 'exit' to quit: ")
		_, err := fmt.Scan(&endpoint)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch endpoint {
		case "create":
			handler.CreateProducts(db)
		case "update":
			handler.UpdateProduct(db)
		case "movement":
			handler.RecordStockMovement(db)
		case "category":
			handler.CreateCategory(db)
		case "location":
			handler.CreateLocation(db)
		case "search":
			handler.SearchItemsByItemCode(db)
		case "ts":
			handler.GetAllTransactions(db)
		case "exit":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Endpoint tidak dikenali. Silakan coba lagi.")
		}
	}
}
