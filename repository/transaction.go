package repository

import (
	"database/sql"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}

func (r *TransactionRepositoryDB) GetAllTransactions() ([]model.InOutProduct, error) {
	rows, err := r.db.Query("SELECT id, product_id, quantity, movement_type, timestamp FROM in_out_product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.InOutProduct
	for rows.Next() {
		var transaction model.InOutProduct
		if err := rows.Scan(&transaction.ID, &transaction.ProductID, &transaction.Quantity, &transaction.MovementType, &transaction.Timestamp); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
