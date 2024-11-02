package service

import (
	"errors"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
)

type ProductService struct {
	RepoProduct repository.ProductRepositoryDB
}

func NewProductService(repo repository.ProductRepositoryDB) *ProductService {
	return &ProductService{RepoProduct: repo}
}

func (cs *ProductService) CreateDataProduct(itemCode, name string, categoryId, locationId, price, stock int) error {
	if name == "" {
		return errors.New("product name cannot be empty")
	}

	product := model.Item{
		ItemCode:   itemCode,
		Name:       name,
		CategoryId: categoryId,
		LocationId: locationId,
		Price:      price,
		Stock:      stock,
	}

	return cs.RepoProduct.Create(&product)
}

func (cs *ProductService) UpdateDataProduct(item model.Item) error {
	if item.Name == "" {
		return errors.New("product name cannot be empty")
	}

	// Assuming this function updates the item in the repository
	return cs.RepoProduct.Update(&item)
}

func (s *ProductService) SearchItems(searchQuery string) ([]model.Item, error) {
	return s.RepoProduct.SearchItems(searchQuery)
}
