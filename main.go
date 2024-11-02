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

	var endpoint string
	fmt.Print("Masukkan endpoint (create/update/movement/category/location/search): ")
	fmt.Scan(&endpoint)
	if err != nil {
		fmt.Println("Error reading input:", err)

	}

	switch endpoint {
	case "create":
		handler.CreateProduct(db)
	case "update":
		handler.UpdateProduct(db)
	case "movement":
		handler.RecordStockMovement(db)
	case "category":
		handler.CreateCategory(db)
	case "location":
		handler.CreateLocation(db)
	case "search":
		handler.SearchItem(db)

	case "exit":
		fmt.Println("Exiting the program.")
		return
	default:
		fmt.Println("Endpoint tidak dikenali. Silakan coba lagi.")
	}

}
