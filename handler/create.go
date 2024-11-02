package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"os"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/service"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/utils"
)

func CreateProducts(db *sql.DB) {
	var products []model.Item
	file, err := os.Open("body.json")
	if err != nil {
		utils.SendErrorResponse("Error reading JSON file: "+err.Error(), nil)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil && err != io.EOF {
		utils.SendErrorResponse("Error decoding JSON: "+err.Error(), nil)
		return
	}

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	for _, product := range products {
		err := productService.CreateDataProduct(product.ItemCode, product.Name, product.CategoryId, product.LocationId, product.Price, product.Stock)
		if err != nil {
			utils.SendErrorResponse("Failed to create product: "+err.Error(), &product)
			return
		}
	}

	responseData := model.Response{
		StatusCode: 200,
		Message:    "All products created successfully",
		Data:       products,
	}

	if err := utils.PrintJSONResponse(responseData); err != nil {
		utils.SendErrorResponse("Error marshaling response: "+err.Error(), nil)
	}
}
