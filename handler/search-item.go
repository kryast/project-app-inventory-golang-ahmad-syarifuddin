package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/service"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/utils"
)

func SearchItem(db *sql.DB) {
	var request struct {
		SearchQuery string `json:"search_query"`
	}

	// Read input from body.json
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

	err = json.Unmarshal(data, &request)
	if err != nil {
		utils.SendErrorResponse("Error decoding JSON: "+err.Error(), nil)
		return
	}

	// Initialize repositories and service
	itemRepo := repository.NewProductRepository(db)
	itemService := service.NewProductService(itemRepo)

	// Search items based on query
	items, err := itemService.SearchItems(request.SearchQuery)
	if err != nil {
		utils.SendErrorResponse("Error searching items: "+err.Error(), nil)
		return
	}

	// Prepare category and location repositories
	categoryRepo := repository.NewCategoryRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	var results []struct {
		ID         int            `json:"id"`
		ItemCode   string         `json:"item_code"`
		Name       string         `json:"name"`
		CategoryId model.Category `json:"category_id"`
		LocationId model.Location `json:"location_id"`
		Price      int            `json:"price"`
		Stock      int            `json:"stock"`
	}

	// Fetch category and location for each item
	for _, item := range items {
		category, err := categoryRepo.FindByID(item.CategoryId)
		if err != nil {
			utils.SendErrorResponse("Error fetching category for ID "+strconv.Itoa(item.CategoryId)+": "+err.Error(), nil)
			continue
		}

		location, err := locationRepo.FindByID(item.LocationId)
		if err != nil {
			utils.SendErrorResponse("Error fetching location for ID "+strconv.Itoa(item.LocationId)+": "+err.Error(), nil)
			continue
		}

		// Append item with details to results
		results = append(results, struct {
			ID         int            `json:"id"`
			ItemCode   string         `json:"item_code"`
			Name       string         `json:"name"`
			CategoryId model.Category `json:"category_id"`
			LocationId model.Location `json:"location_id"`
			Price      int            `json:"price"`
			Stock      int            `json:"stock"`
		}{
			ID:         item.ID,
			ItemCode:   item.ItemCode,
			Name:       item.Name,
			CategoryId: *category,
			LocationId: *location,
			Price:      item.Price,
			Stock:      item.Stock,
		})
	}

	// Convert response to model.Response
	responseData := model.Response{
		StatusCode: 200,
		Message:    "Items found",
		Data:       results,
	}

	if err := utils.PrintJSONResponse(responseData); err != nil {
		utils.SendErrorResponse("Error marshaling response to JSON: "+err.Error(), nil)
	}
}
