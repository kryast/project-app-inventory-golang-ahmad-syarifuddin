package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"os"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/utils"
)

func CreateCategory(db *sql.DB) {
	var category model.Category

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

	err = json.Unmarshal(data, &category)
	if err != nil {
		utils.SendErrorResponse("Error decoding JSON: "+err.Error(), nil)
		return
	}

	repo := repository.NewCategoryRepository(db)
	if err := repo.Create(&category); err != nil {
		utils.SendErrorResponse("Error while creating category: "+err.Error(), nil)
		return
	}

	response := model.Response{
		StatusCode: 200,
		Message:    "Category created successfully",
		Data: struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   category.ID,
			Name: category.Name,
		},
	}

	if err := utils.PrintJSONResponse(response); err != nil {
		utils.SendErrorResponse("Error marshaling response to JSON: "+err.Error(), nil)
	}
}
