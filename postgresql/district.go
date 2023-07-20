package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
)

type DistrictStore struct {
	db *sql.DB
}

func (store DistrictStore) Save(ctx context.Context, district models.District) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO districts (code, name, population, region_id) VALUES ($1, $2, $3, $4)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %s", err)
	}

	if _, err := stmt.Exec(district.Code, district.Name, district.Population, district.RegionId); err != nil {
		return fmt.Errorf("error saving district: %s", err)
	}

	return nil
}

func (store DistrictStore) Find(id int) (models.District, error) {
	row := store.db.QueryRow("SELECT * FROM districts WHERE id = $1", id)
	district, err := newDistrict(row)

	if err == nil {
		return district, nil
	}

	if err == sql.ErrNoRows {
		return district, fmt.Errorf("district with id = %d not found: %s", id, err)
	}

	return district, fmt.Errorf("error executing query: %s", err)
}

func (store DistrictStore) FindByCode(code string) (models.District, error) {
	row := store.db.QueryRow("SELECT * FROM districts WHERE code = $1", code)
	district, err := newDistrict(row)

	if err == nil {
		return district, nil
	}

	if err == sql.ErrNoRows {
		return district, fmt.Errorf("district with code = %s not found: %s", code, err)
	}

	return district, fmt.Errorf("error executing query: %s", err)
}

func (store DistrictStore) FindByName(name string) (models.District, error) {
	row := store.db.QueryRow("SELECT * FROM districts WHERE name = $1", name)
	district, err := newDistrict(row)

	if err == nil {
		return district, nil
	}

	if err == sql.ErrNoRows {
		return district, fmt.Errorf("district with name = %s not found: %s", name, err)
	}

	return district, fmt.Errorf("error executing query: %s", err)
}

func newDistrict(row *sql.Row) (models.District, error) {
	var district models.District

	err := row.Scan(
		&district.Id,
		&district.Code,
		&district.Name,
		&district.Population,
		&district.RegionId,
	)

	return district, err
}
