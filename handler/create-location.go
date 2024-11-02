package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/utils"
)

func CreateLocation(db *sql.DB) {
	var location model.Location

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

	// Log data mentah yang dibaca dari file
	fmt.Println("Raw JSON data:", string(data))

	err = json.Unmarshal(data, &location)
	if err != nil {
		utils.SendErrorResponse("Error decoding JSON: "+err.Error(), nil)
		return
	}

	// Log nilai location setelah unmarshal
	fmt.Printf("Location after unmarshal: %+v\n", location)

	repo := repository.NewLocationRepository(db)
	if err := repo.Create(&location); err != nil {
		utils.SendErrorResponse("Error while creating location: "+err.Error(), nil)
		return
	}

	response := model.Response{
		StatusCode: 200,
		Message:    "Location created successfully",
		Data: struct {
			ID           int    `json:"id"`
			Address      string `json:"address"`
			City         string `json:"city"`
			Province     string `json:"province"`
			ItemPosition string `json:"item_position"`
		}{
			ID:           location.ID,
			Address:      location.Address,
			City:         location.City,
			Province:     location.Province,
			ItemPosition: location.ItemPosition,
		},
	}

	if err := utils.PrintJSONResponse(response); err != nil {
		utils.SendErrorResponse("Error marshaling response to JSON: "+err.Error(), nil)
	}
}
