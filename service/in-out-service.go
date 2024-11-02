package service

import (
	"time"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
)

type InOutProductService struct {
	RepoInOutProduct repository.InOutProductRepositoryDB
	RepoItem         repository.ProductRepositoryDB
}

func NewInOutProductService(repo repository.InOutProductRepositoryDB, itemRepo repository.ProductRepositoryDB) *InOutProductService {
	return &InOutProductService{RepoInOutProduct: repo, RepoItem: itemRepo}
}

func (s *InOutProductService) RecordMovement(movement *model.InOutProduct) error {
	if movement.Timestamp.IsZero() {
		movement.Timestamp = time.Now()
	}
	return s.RepoInOutProduct.Create(movement)
}

func (s *InOutProductService) GetItemByID(id int) (*model.Item, error) {
	return s.RepoItem.FindByID(id)
}
