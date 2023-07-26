package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
)

type CityStore struct {
	db *sql.DB
}

func NewCityStore(db *sql.DB) *CityStore {
	return &CityStore{db: db}
}

func (store CityStore) Save(ctx context.Context, city models.City) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO cities (code, name, city_class, income_class, population, province_code, district_code) VALUES ($1, $2, $3, $4, $5, $6, $7)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %s", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(
		city.Code,
		city.Name,
		city.CityClass,
		city.IncomeClass,
		city.Population,
		city.ProvinceCode,
		city.DistrictCode,
	); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (store CityStore) Find(id int) (models.City, error) {
	row := store.db.QueryRow("SELECT * FROM cities WHERE cities.id = $1", id)
	city, err := newCity(row)

	if err == nil {
		return city, nil
	}

	if err == sql.ErrNoRows {
		return city, fmt.Errorf("city with id = %d not found: %w", id, err)
	}

	return city, fmt.Errorf("error executing query: %w", err)
}

func (store CityStore) FindByCode(code string) (models.City, error) {
	row := store.db.QueryRow("SELECT * FROM cities WHERE code = $1", code)
	city, err := newCity(row)

	if err == nil {
		return city, nil
	}

	if err == sql.ErrNoRows {
		return city, fmt.Errorf("city with code = %s not found: %w", code, err)
	}

	return city, fmt.Errorf("error executing query: %w", err)
}

func (store CityStore) FindByName(name string) (models.City, error) {
	row := store.db.QueryRow("SELECT * FROM cities WHERE name = $1", name)
	city, err := newCity(row)

	if err == nil {
		return city, nil
	}

	if err == sql.ErrNoRows {
		return city, fmt.Errorf("city with name = %s not found: %w", name, err)
	}

	return city, fmt.Errorf("error executing query: %w", err)
}

func newCity(row *sql.Row) (models.City, error) {
	var city models.City

	err := row.Scan(
		&city.Id,
		&city.Code,
		&city.Name,
		&city.CityClass,
		&city.IncomeClass,
		&city.Population,
		&city.ProvinceCode,
		&city.DistrictCode,
	)

	if err != nil {
		return city, err
	}

	return city, nil
}
