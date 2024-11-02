package repository

import (
	"database/sql"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

type ProductRepositoryDB struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryDB {
	return ProductRepositoryDB{DB: db}
}

func (r *ProductRepositoryDB) Create(item *model.Item) error {
	query := `INSERT INTO item (item_code, name, category_id, location_id, price, stock) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.DB.QueryRow(query, item.ItemCode, item.Name, item.CategoryId, item.LocationId, item.Price, item.Stock).Scan(&item.ID)
	if err != nil {
		return err
	}

	return nil
}
