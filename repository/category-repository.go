package repository

import (
	"database/sql"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

type CategoryRepositoryDB interface {
	FindByID(id int) (*model.Category, error)
	Create(category *model.Category) error
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepositoryDB {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindByID(id int) (*model.Category, error) {
	query := `SELECT id, category_name FROM category WHERE id = $1` // Pastikan nama tabel kategori sesuai
	row := r.db.QueryRow(query, id)

	var category model.Category
	if err := row.Scan(&category.ID, &category.Name); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Create(category *model.Category) error {
	query := `INSERT INTO category (category_name) VALUES ($1) RETURNING id`
	return r.db.QueryRow(query, category.Name).Scan(&category.ID)
}
