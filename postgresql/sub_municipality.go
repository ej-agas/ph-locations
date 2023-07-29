package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
)

type SubMunicipalityStore struct {
	db *sql.DB
}

func NewSubMunicipalityStore(db *sql.DB) *SubMunicipalityStore {
	return &SubMunicipalityStore{db: db}
}

func (store SubMunicipalityStore) Save(ctx context.Context, subMunicipality models.SubMunicipality) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO sub_municipalities (code, name, population, city_code) VALUES ($1, $2, $3, $4)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %w", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(
		subMunicipality.Code,
		subMunicipality.Name,
		subMunicipality.Population,
		subMunicipality.CityCode,
	); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (store SubMunicipalityStore) Find(id int) (models.SubMunicipality, error) {
	row := store.db.QueryRow("SELECT * FROM sub_municipalities WHERE id = $1", id)
	subMunicipality, err := newSubMunicipality(row)

	if err == nil {
		return subMunicipality, nil
	}

	if err == sql.ErrNoRows {
		return subMunicipality, fmt.Errorf("SubMunicipality with id = %d not found: %w", id, err)
	}

	return subMunicipality, fmt.Errorf("error executing query: %w", err)
}

func (store SubMunicipalityStore) FindByCode(code string) (models.SubMunicipality, error) {
	row := store.db.QueryRow("SELECT * FROM sub_municipalities WHERE code = $1", code)
	subMunicipality, err := newSubMunicipality(row)

	if err == nil {
		return subMunicipality, nil
	}

	if err == sql.ErrNoRows {
		return subMunicipality, fmt.Errorf("SubMunicipality with code = %s not found: %w", code, err)
	}

	return subMunicipality, fmt.Errorf("error executing query: %w", err)
}

func (store SubMunicipalityStore) FindByName(name string) (models.SubMunicipality, error) {
	row := store.db.QueryRow("SELECT * FROM sub_municipalities WHERE name = $1", name)
	subMunicipality, err := newSubMunicipality(row)

	if err == nil {
		return subMunicipality, nil
	}

	if err == sql.ErrNoRows {
		return subMunicipality, fmt.Errorf("SubMunicipality with name = %s not found: %w", name, err)
	}

	return subMunicipality, fmt.Errorf("error executing query: %w", err)
}

func (store SubMunicipalityStore) List(opts stores.SearchOpts) (stores.Collection[models.SubMunicipality], error) {
	collection := stores.Collection[models.SubMunicipality]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from sub_municipalities").Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM sub_municipalities ORDER BY $1 LIMIT $2 OFFSET $3",
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	subMunicipalities, err := newSubMunicipalities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = subMunicipalities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store SubMunicipalityStore) ListByCityCode(code string, opts stores.SearchOpts) (stores.Collection[models.SubMunicipality], error) {
	collection := stores.Collection[models.SubMunicipality]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from sub_municipalities WHERE city_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM sub_municipalities WHERE city_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	subMunicipalities, err := newSubMunicipalities(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = subMunicipalities
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newSubMunicipalities(rows *sql.Rows, count int) ([]models.SubMunicipality, error) {
	subMunicipalities := make([]models.SubMunicipality, 0, count)

	for rows.Next() {
		var s models.SubMunicipality

		if err := rows.Scan(
			&s.Id,
			&s.Code,
			&s.Name,
			&s.Population,
			&s.CityCode,
		); err != nil {
			return subMunicipalities, err
		}
		subMunicipalities = append(subMunicipalities, s)
	}

	return subMunicipalities, nil
}

func newSubMunicipality(row *sql.Row) (models.SubMunicipality, error) {
	var subMun models.SubMunicipality

	err := row.Scan(
		&subMun.Id,
		&subMun.Code,
		&subMun.Name,
		&subMun.Population,
		&subMun.CityCode,
	)

	return subMun, err
}
