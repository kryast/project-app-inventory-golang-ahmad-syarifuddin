package service

import (
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepositoryDB
}

func NewTransactionService(repo *repository.TransactionRepositoryDB) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) GetAllTransactions() ([]model.InOutProduct, error) {
	return s.repo.GetAllTransactions()
}
