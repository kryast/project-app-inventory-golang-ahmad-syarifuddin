package utils

import (
	"encoding/json"
	"fmt"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

// SendErrorResponse sends a JSON error response
func SendErrorResponse(message string, data interface{}) {
	response := model.Response{
		StatusCode: 400,
		Message:    message,
		Data:       data,
	}
	if err := PrintJSONResponse(response); err != nil {
		fmt.Println("Error marshaling error response:", err)
	}
}

// PrintJSONResponse prints the JSON response to the console (or your response writer)
func PrintJSONResponse(response model.Response) error {
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData)) // Replace with your response writer in a real server
	return nil
}
