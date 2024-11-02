package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/service"
)

func CreateProduct(db *sql.DB) {
	// Input dari file JSON
	var product model.Item
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&product)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding JSON: ", err)
		return
	}

	// Proses
	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	err = productService.CreateDataProduct(product.ItemCode, product.Name, product.CategoryId, product.LocationId, product.Price, product.Stock)
	if err != nil {
		fmt.Println("Error while creating product: ", err)
		return
	}

	// Output yang diinginkan
	response := model.Response{
		StatusCode: 200,
		Message:    "create success",
		Data:       product,
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling response: ", err)
		return
	}

	fmt.Println(string(jsonData))
}
