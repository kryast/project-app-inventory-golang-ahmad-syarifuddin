package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/service"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/utils"
)

func RecordStockMovement(db *sql.DB) {
	var inOut model.InOutProduct

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

	err = json.Unmarshal(data, &inOut)
	if err != nil {
		utils.SendErrorResponse("Error decoding JSON: "+err.Error(), nil)
		return
	}

	if inOut.Timestamp.IsZero() {
		inOut.Timestamp = time.Now()
	}

	inOutRepo := repository.NewInOutProductRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	inOutProductService := service.NewInOutProductService(inOutRepo, productRepo)

	if err := inOutProductService.RecordMovement(&inOut); err != nil {
		utils.SendErrorResponse("Error while recording stock movement: "+err.Error(), nil)
		return
	}

	product, err := productRepo.FindByID(inOut.ProductID)
	if err != nil {
		utils.SendErrorResponse("Error fetching product details: "+err.Error(), nil)
		return
	}

	category, err := categoryRepo.FindByID(product.CategoryId)
	if err != nil {
		utils.SendErrorResponse("Error fetching category details: "+err.Error(), nil)
		return
	}

	location, err := locationRepo.FindByID(product.LocationId)
	if err != nil {
		utils.SendErrorResponse("Error fetching location details: "+err.Error(), nil)
		return
	}

	formattedTimestamp := inOut.Timestamp.Format("2006-01-02 15:04:05")

	responseData := model.Response{
		StatusCode: 200,
		Message:    "Stock movement recorded successfully",
		Data: struct {
			ID           int            `json:"id"`
			Product      model.Item     `json:"product"`
			Category     model.Category `json:"category"`
			Location     model.Location `json:"location"`
			Quantity     int            `json:"quantity"`
			MovementType string         `json:"movement_type"`
			Timestamp    string         `json:"timestamp"`
		}{
			ID:           inOut.ID,
			Product:      *product,
			Category:     *category,
			Location:     *location,
			Quantity:     inOut.Quantity,
			MovementType: inOut.MovementType,
			Timestamp:    formattedTimestamp,
		},
	}

	if err := utils.PrintJSONResponse(responseData); err != nil {
		utils.SendErrorResponse("Error marshaling response to JSON: "+err.Error(), nil)
	}
}
