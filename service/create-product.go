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
		return errors.New("username tidak boleh kosong")
	}

	product := model.Item{
		ItemCode:   item_code,
		Name:       name,
		CategoryId: category_id,
		LocationId: location_id,
		Price:      price,
		Stock:      stock,
	}

	err := cs.RepoProduct.Create(&product)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("berhasil input data Admin dengan id ", product.ID)

	return nil
}
