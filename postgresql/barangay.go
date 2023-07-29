package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
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
		"INSERT INTO barangays (code, name, urban_rural, population, city_code, municipality_code, sub_municipality_code, special_government_unit_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
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
		barangay.CityCode,
		barangay.MunicipalityCode,
		barangay.SubMunicipalityCode,
		barangay.SpecialGovernmentUnitCode,
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

	if err == nil {
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

func (store BarangayStore) ListByCityCode(code string, opts stores.SearchOpts) (stores.Collection[models.Barangay], error) {
	collection := stores.Collection[models.Barangay]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from barangays WHERE city_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM barangays WHERE city_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	barangays, err := newBarangays(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = barangays
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store BarangayStore) ListByMunicipalityCode(code string, opts stores.SearchOpts) (stores.Collection[models.Barangay], error) {
	collection := stores.Collection[models.Barangay]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from barangays WHERE municipality_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM barangays WHERE municipality_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	barangays, err := newBarangays(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = barangays
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store BarangayStore) ListBySubMunicipalityCode(code string, opts stores.SearchOpts) (stores.Collection[models.Barangay], error) {
	collection := stores.Collection[models.Barangay]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from barangays WHERE sub_municipality_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM barangays WHERE sub_municipality_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	barangays, err := newBarangays(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = barangays
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store BarangayStore) ListBySpecialGovernmentUnitCode(code string, opts stores.SearchOpts) (stores.Collection[models.Barangay], error) {
	collection := stores.Collection[models.Barangay]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from barangays WHERE special_government_unit_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM barangays WHERE special_government_unit_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	barangays, err := newBarangays(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = barangays
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newBarangays(rows *sql.Rows, count int) ([]models.Barangay, error) {
	barangays := make([]models.Barangay, 0, count)

	for rows.Next() {
		var barangay models.Barangay

		if err := rows.Scan(
			&barangay.Id,
			&barangay.Code,
			&barangay.Name,
			&barangay.UrbanRural,
			&barangay.Population,
			&barangay.CityCode,
			&barangay.MunicipalityCode,
			&barangay.SubMunicipalityCode,
			&barangay.SpecialGovernmentUnitCode,
		); err != nil {
			return barangays, err
		}
		barangays = append(barangays, barangay)
	}

	return barangays, nil
}

func newBarangay(row *sql.Row) (models.Barangay, error) {
	var barangay models.Barangay

	err := row.Scan(
		&barangay.Id,
		&barangay.Code,
		&barangay.Name,
		&barangay.UrbanRural,
		&barangay.Population,
		&barangay.CityCode,
		&barangay.MunicipalityCode,
		&barangay.SubMunicipalityCode,
		&barangay.SpecialGovernmentUnitCode,
	)

	return barangay, err
}
