package repository

import (
	"database/sql"

	"github.com/kryast/project-app-inventory-golang-ahmad-syarifuddin/model"
)

type LocationRepositoryDB interface {
	FindByID(id int) (*model.Location, error)
	Create(location *model.Location) error
}

type LocationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) LocationRepositoryDB {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) FindByID(id int) (*model.Location, error) {
	query := `SELECT id, address, city, province, item_position FROM location WHERE id = $1` // Pastikan nama tabel lokasi sesuai
	row := r.db.QueryRow(query, id)

	var location model.Location
	if err := row.Scan(&location.ID, &location.Address, &location.City, &location.Province, &location.ItemPosition); err != nil {
		return nil, err
	}

	return &location, nil
}

func (r *LocationRepository) Create(location *model.Location) error {
	query := `INSERT INTO location (address, city, province, item_position) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, location.Address, location.City, location.Province, location.ItemPosition).Scan(&location.ID)
}
