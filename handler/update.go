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

func UpdateProduct(db *sql.DB) {
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

	// Check if CategoryId and LocationId are valid
	categoryRepo := repository.NewCategoryRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	if _, err := categoryRepo.FindByID(item.CategoryId); err != nil {
		utils.SendErrorResponse("Invalid CategoryId: "+err.Error(), nil)
		return
	}

	if _, err := locationRepo.FindByID(item.LocationId); err != nil {
		utils.SendErrorResponse("Invalid LocationId: "+err.Error(), nil)
		return
	}

	// Update the item
	repo := repository.NewProductRepository(db)
	itemService := service.NewProductService(repo)

	if err := itemService.UpdateDataProduct(item); err != nil {
		utils.SendErrorResponse("Error while updating item: "+err.Error(), nil)
		return
	}

	// Prepare and send response
	response := model.Response{
		StatusCode: 200,
		Message:    "Update successful",
		Data:       item,
	}

	if err := utils.PrintJSONResponse(response); err != nil {
		utils.SendErrorResponse("Error marshaling response to JSON: "+err.Error(), nil)
	}
}
