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
	fmt.Print("masukkan enpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {

	case "create":
		handler.CreateProduct(db)
	}

	// repo := repository.NewProductRepository(db)
	// adminService := service.NewProductService(repo)

	// adminService.
}
