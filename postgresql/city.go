package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
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

func (store CityStore) ListByProvinceCode(code string, opts stores.SearchOpts) (stores.Collection[models.City], error) {
	collection := stores.Collection[models.City]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from cities WHERE province_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM cities WHERE province_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	cities, err := newCities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = cities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store CityStore) ListByDistrictCode(code string, opts stores.SearchOpts) (stores.Collection[models.City], error) {
	collection := stores.Collection[models.City]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from cities WHERE district_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM cities WHERE district_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	cities, err := newCities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = cities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newCities(rows *sql.Rows, count int) ([]models.City, error) {
	cities := make([]models.City, 0, count)

	for rows.Next() {
		var city models.City

		if err := rows.Scan(
			&city.Id,
			&city.Code,
			&city.Name,
			&city.CityClass,
			&city.IncomeClass,
			&city.Population,
			&city.ProvinceCode,
			&city.DistrictCode,
		); err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}

	return cities, nil
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
