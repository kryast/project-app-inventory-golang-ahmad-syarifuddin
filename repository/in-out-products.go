package repository

import (
	"database/sql"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

type InOutProductRepositoryDB interface {
	Create(movement *model.InOutProduct) error
	FindByProductID(productID int) ([]model.InOutProduct, error)
}

type InOutProductRepository struct {
	db *sql.DB
}

func NewInOutProductRepository(db *sql.DB) InOutProductRepositoryDB {
	return &InOutProductRepository{db: db}
}

func (r *InOutProductRepository) Create(movement *model.InOutProduct) error {
	query := `INSERT INTO in_out_product (product_id, quantity, movement_type, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, movement.ProductID, movement.Quantity, movement.MovementType, movement.Timestamp)
	return err
}

func (r *InOutProductRepository) FindByProductID(productID int) ([]model.InOutProduct, error) {
	query := `SELECT id, product_id, quantity, movement_type, timestamp FROM in_out_product WHERE product_id = $1`
	rows, err := r.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []model.InOutProduct
	for rows.Next() {
		var movement model.InOutProduct
		if err := rows.Scan(&movement.ID, &movement.ProductID, &movement.Quantity, &movement.MovementType, &movement.Timestamp); err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}

	return movements, nil
}
