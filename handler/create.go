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

func CreateProduct(db *sql.DB) {
	var item model.Item

	file, err := os.Open("body.json")
	if err != nil {
		utils.SendErrorResponse("Error opening JSON file: "+err.Error(), nil)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		utils.SendErrorResponse("Error reading JSON file: "+err.Error(), nil)
		return
	}

	err = json.Unmarshal(data, &item)
	if err != nil {
		utils.SendErrorResponse("Error decoding JSON: "+err.Error(), nil)
		return
	}

	// Create the item
	repo := repository.NewProductRepository(db)
	itemService := service.NewProductService(repo)

	if err := itemService.CreateDataProduct(item.ItemCode, item.Name, item.CategoryId, item.LocationId, item.Price, item.Stock); err != nil {
		utils.SendErrorResponse("Error while creating item: "+err.Error(), nil)
		return
	}

	// Prepare and send response
	response := model.Response{
		StatusCode: 200,
		Message:    "Item created successfully",
		Data:       item,
	}

	if err := utils.PrintJSONResponse(response); err != nil {
		utils.SendErrorResponse("Error marshaling response to JSON: "+err.Error(), nil)
	}
}
