package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
)

type DistrictStore struct {
	db *sql.DB
}

func NewDistrictStore(connection *sql.DB) *DistrictStore {
	return &DistrictStore{db: connection}
}

func (store DistrictStore) Save(ctx context.Context, district models.District) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO districts (code, name, population, region_code) VALUES ($1, $2, $3, $4)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %w", err)
	}

	if _, err := stmt.Exec(district.Code, district.Name, district.Population, district.RegionCode); err != nil {
		return fmt.Errorf("error saving district: %w", err)
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
		return district, fmt.Errorf("district with id = %d not found: %w", id, err)
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
		return district, fmt.Errorf("district with code = %s not found: %w", code, err)
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
		return district, fmt.Errorf("district with name = %s not found: %w", name, err)
	}

	return district, fmt.Errorf("error executing query: %w", err)
}

func (store DistrictStore) List(opts stores.SearchOpts) (stores.Collection[models.District], error) {
	collection := stores.Collection[models.District]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from districts").Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	q := fmt.Sprintf("SELECT * FROM districts ORDER BY %s %s LIMIT $1 OFFSET $2", opts.Order, opts.Sort)
	rows, err := store.db.Query(q, opts.Limit, offset)

	if err != nil {
		return collection, err
	}

	districts, err := newDistricts(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = districts
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store DistrictStore) ListByRegionCode(code string, opts stores.SearchOpts) (stores.Collection[models.District], error) {
	collection := stores.Collection[models.District]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from districts WHERE region_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	q := fmt.Sprintf("SELECT * FROM districts WHERE region_code = $1 ORDER BY %s %s LIMIT $2 OFFSET $3", opts.Order, opts.Sort)
	rows, err := store.db.Query(q, code, opts.Limit, offset)

	if err != nil {
		return collection, err
	}

	districts, err := newDistricts(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = districts
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newDistricts(rows *sql.Rows, count int) ([]models.District, error) {
	districts := make([]models.District, 0, count)

	for rows.Next() {
		var d models.District

		if err := rows.Scan(
			&d.Id,
			&d.Code,
			&d.Name,
			&d.Population,
			&d.RegionCode,
		); err != nil {
			return districts, err
		}
		districts = append(districts, d)
	}

	return districts, nil
}

func newDistrict(row *sql.Row) (models.District, error) {
	var district models.District

	err := row.Scan(
		&district.Id,
		&district.Code,
		&district.Name,
		&district.Population,
		&district.RegionCode,
	)

	return district, err
}
