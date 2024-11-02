package repository

import (
	"database/sql"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

type ProductRepositoryDB interface {
	Create(item *model.Item) error
	Update(item *model.Item) error
	FindByID(id int) (*model.Item, error)
	SearchItems(searchQuery string) ([]model.Item, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryDB {
	return &productRepository{db: db}
}

func (r *productRepository) Create(item *model.Item) error {
	// Implementation for inserting a new product into the database
	_, err := r.db.Exec("INSERT INTO item (item_code, name, category_id, location_id, price, stock) VALUES ($1, $2, $3, $4, $5, $6)",
		item.ItemCode, item.Name, item.CategoryId, item.LocationId, item.Price, item.Stock)
	return err
}

func (r *productRepository) Update(item *model.Item) error {
	// Implementation for updating an existing product
	_, err := r.db.Exec("UPDATE item SET item_code = $1, name = $2, category_id = $3, location_id = $4, price = $5, stock = $6 WHERE id = $7",
		item.ItemCode, item.Name, item.CategoryId, item.LocationId, item.Price, item.Stock, item.ID)
	return err
}

func (r *productRepository) FindByID(id int) (*model.Item, error) {
	query := `SELECT id, item_code, name, category_id, location_id, price, stock FROM item WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var item model.Item
	if err := row.Scan(&item.ID, &item.ItemCode, &item.Name, &item.CategoryId, &item.LocationId, &item.Price, &item.Stock); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *productRepository) SearchItems(searchQuery string) ([]model.Item, error) {
	query := `
		SELECT id, item_code, name, category_id, location_id, price, stock
		FROM item
		WHERE name ILIKE $1 OR item_code ILIKE $1 OR category_id IN (SELECT id FROM category WHERE name ILIKE $1)
	`
	rows, err := r.db.Query(query, "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.ItemCode, &item.Name, &item.CategoryId, &item.LocationId, &item.Price, &item.Stock); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
