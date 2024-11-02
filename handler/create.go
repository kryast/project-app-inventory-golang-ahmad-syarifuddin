package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/service"
)

const inputFileName = "body.json"

func CreateProduct(db *sql.DB) {
	var product model.Item

	if err := readJSONFile(inputFileName, &product); err != nil {
		sendErrorResponse("Error reading JSON file: " + err.Error())
		return
	}

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)

	if err := productService.CreateDataProduct(product.ItemCode, product.Name, product.CategoryId, product.LocationId, product.Price, product.Stock); err != nil {

		sendErrorResponse("Error while creating product: " + err.Error())
		return
	} else {

		response := model.Response{
			StatusCode: 200,
			Message:    "create success",
			Data:       product,
		}

		if err := printJSONResponse(response); err != nil {
			sendErrorResponse("Error marshaling response: " + err.Error())
		}
	}
}

func readJSONFile(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}

func printJSONResponse(response model.Response) error {
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}

func sendErrorResponse(message string) {
	response := model.Response{
		StatusCode: 400,
		Message:    message,
		Data:       model.Item{},
	}
	if err := printJSONResponse(response); err != nil {
		fmt.Println("Error marshaling error response:", err)
	}
}
