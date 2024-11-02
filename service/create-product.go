package service

import (
	"errors"
	"fmt"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
)

type ProductService struct {
	RepoProduct repository.ProductRepositoryDB
}

func NewProductService(repo repository.ProductRepositoryDB) *ProductService {
	return &ProductService{RepoProduct: repo}
}

func (cs *ProductService) CreateDataProduct(item_code string, name string, category_id int, location_id int, price int, stock int) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	product := model.Item{
		ItemCode:   item_code,
		Name:       name,
		CategoryId: category_id,
		LocationId: location_id,
		Price:      price,
		Stock:      stock,
	}

	// Attempt to create the product
	err := cs.RepoProduct.Create(&product)
	if err != nil {
		// Log the error
		fmt.Println("Error while creating product:", err)
		return errors.New("failed to create product") // Return a generic error message
	}

	// Print success message if creation is successful
	fmt.Println("Successfully created product with ID:", product.ID)
	return nil
}
