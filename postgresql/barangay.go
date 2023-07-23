package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
)

type BarangayStore struct {
	db *sql.DB
}

func NewBarangayStore(db *sql.DB) *BarangayStore {
	return &BarangayStore{db: db}
}

func (store BarangayStore) Save(ctx context.Context, barangay models.Barangay) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO barangays (code, name, urban_rural, population, city_id, municipality_id, sub_municipality_id, special_government_unit_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %w", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(
		barangay.Code,
		barangay.Name,
		barangay.UrbanRural,
		barangay.Population,
		barangay.CityId,
		barangay.MunicipalityId,
		barangay.SubMunicipalityId,
		barangay.SpecialGovernmentUnitId,
	); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (store BarangayStore) Find(id int) (models.Barangay, error) {
	row := store.db.QueryRow("SELECT * FROM barangays WHERE id = $1", id)
	barangay, err := newBarangay(row)

	if err == nil {
		return barangay, nil
	}

	if err == sql.ErrNoRows {
		return barangay, fmt.Errorf("barangay with id = %d not found: %w", id, err)
	}

	return barangay, fmt.Errorf("error executing query: %w", err)
}

func (store BarangayStore) FindByCode(code string) (models.Barangay, error) {
	row := store.db.QueryRow("SELECT * FROM barangays WHERE code = $1", code)
	barangay, err := newBarangay(row)

	if err != nil {
		return barangay, nil
	}

	if err == sql.ErrNoRows {
		return barangay, fmt.Errorf("barangay with code = %s not found: %w", code, err)
	}

	return barangay, fmt.Errorf("error executing query: %w", err)
}

func (store BarangayStore) FindByName(name string) (models.Barangay, error) {
	row := store.db.QueryRow("SELECT * FROM barangays WHERE name = $1", name)
	barangay, err := newBarangay(row)

	if err != nil {
		return barangay, nil
	}

	if err == sql.ErrNoRows {
		return barangay, fmt.Errorf("barangay with name = %s not found: %w", name, err)
	}

	return barangay, fmt.Errorf("error executing query: %w", err)
}

func newBarangay(row *sql.Row) (models.Barangay, error) {
	var barangay models.Barangay

	err := row.Scan(
		&barangay.Id,
		&barangay.Code,
		&barangay.Name,
		&barangay.UrbanRural,
		&barangay.Population,
		&barangay.CityId,
		&barangay.MunicipalityId,
		&barangay.SubMunicipalityId,
		&barangay.SpecialGovernmentUnitId,
	)

	return barangay, err
}
