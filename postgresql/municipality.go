package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
)

type MunicipalityStore struct {
	db *sql.DB
}

func NewMunicipalityStore(connection *sql.DB) *MunicipalityStore {
	return &MunicipalityStore{db: connection}
}

func (store MunicipalityStore) Save(ctx context.Context, municipality models.Municipality) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO municipalities (code, name, income_class, population, province_code, district_code) VALUES ($1, $2, $3, $4, $5, $6)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %s", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(
		municipality.Code,
		municipality.Name,
		municipality.IncomeClass,
		municipality.Population,
		municipality.ProvinceCode,
		municipality.DistrictCode,
	); err != nil {
		return fmt.Errorf("error executing query: %s", err)
	}

	return nil
}

func (store MunicipalityStore) Find(id int) (models.Municipality, error) {
	row := store.db.QueryRow("SELECT * FROM municipalities WHERE municipalities.id = $1", id)
	municipality, err := newMunicipality(row)

	if err == nil {
		return municipality, nil
	}

	if err == sql.ErrNoRows {
		return municipality, fmt.Errorf("municipality with id = %d not found: %s", id, err)
	}

	return municipality, fmt.Errorf("error executing query: %s", err)
}

func (store MunicipalityStore) FindByCode(code string) (models.Municipality, error) {
	row := store.db.QueryRow("SELECT * FROM municipalities WHERE code = $1", code)
	municipality, err := newMunicipality(row)

	if err == nil {
		return municipality, nil
	}

	if err == sql.ErrNoRows {
		return municipality, fmt.Errorf("municipality with code = %s not found: %s", code, err)
	}

	return municipality, fmt.Errorf("error executing query: %s", err)
}

func (store MunicipalityStore) FindByName(name string) (models.Municipality, error) {
	row := store.db.QueryRow("SELECT * FROM municipalities WHERE name = $1", name)
	municipality, err := newMunicipality(row)

	if err == nil {
		return municipality, nil
	}

	if err == sql.ErrNoRows {
		return municipality, fmt.Errorf("municipality with name = %s not found: %s", name, err)
	}

	return municipality, fmt.Errorf("error executing query: %s", err)
}

func (store MunicipalityStore) List(opts stores.SearchOpts) (stores.Collection[models.Municipality], error) {
	collection := stores.Collection[models.Municipality]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(*) from municipalities").Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	q := fmt.Sprintf("SELECT * FROM municipalities ORDER BY %s %s LIMIT $2 OFFSET $3", opts.Order, opts.Sort)
	rows, err := store.db.Query(q, opts.Limit, offset)

	if err != nil {
		return collection, err
	}

	municipalities, err := newMunicipalities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		Total:       int(totalRows),
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = municipalities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store MunicipalityStore) ListByProvinceCode(code string, opts stores.SearchOpts) (stores.Collection[models.Municipality], error) {
	collection := stores.Collection[models.Municipality]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(*) from municipalities WHERE province_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	q := fmt.Sprintf("SELECT * FROM municipalities WHERE province_code = $1 ORDER BY %s %s LIMIT $2 OFFSET $3", opts.Order, opts.Sort)
	rows, err := store.db.Query(q, code, opts.Limit, offset)

	if err != nil {
		return collection, err
	}

	municipalities, err := newMunicipalities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		Total:       int(totalRows),
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = municipalities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store MunicipalityStore) ListByDistrictCode(code string, opts stores.SearchOpts) (stores.Collection[models.Municipality], error) {
	collection := stores.Collection[models.Municipality]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(*) from municipalities WHERE district_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	q := fmt.Sprintf("SELECT * FROM municipalities WHERE district_code = $1 ORDER BY %s %s LIMIT $2 OFFSET $3", opts.Order, opts.Sort)
	rows, err := store.db.Query(q, code, opts.Limit, offset)

	if err != nil {
		return collection, err
	}

	cities, err := newMunicipalities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		Total:       int(totalRows),
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = cities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newMunicipalities(rows *sql.Rows, count int) ([]models.Municipality, error) {
	municipalities := make([]models.Municipality, 0, count)

	for rows.Next() {
		var m models.Municipality

		if err := rows.Scan(
			&m.Id,
			&m.Code,
			&m.Name,
			&m.IncomeClass,
			&m.Population,
			&m.ProvinceCode,
			&m.DistrictCode,
		); err != nil {
			return municipalities, err
		}
		municipalities = append(municipalities, m)
	}

	return municipalities, nil
}

func newMunicipality(row *sql.Row) (models.Municipality, error) {
	var m models.Municipality

	err := row.Scan(
		&m.Id,
		&m.Code,
		&m.Name,
		&m.IncomeClass,
		&m.Population,
		&m.ProvinceCode,
		&m.DistrictCode,
	)

	if err != nil {
		return m, err
	}

	return m, nil
}
