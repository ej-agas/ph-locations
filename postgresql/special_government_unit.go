package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
)

type SpecialGovernmentUnit struct {
	db *sql.DB
}

func NewSpecialGovernmentUnit(db *sql.DB) *SpecialGovernmentUnit {
	return &SpecialGovernmentUnit{db: db}
}

func (store SpecialGovernmentUnit) Save(ctx context.Context, sgu models.SpecialGovernmentUnit) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO special_government_units (code, name, province_code) VALUES ($1, $2, $3)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %w", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(sgu.Code, sgu.Name, sgu.ProvinceCode); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (store SpecialGovernmentUnit) Find(id int) (models.SpecialGovernmentUnit, error) {
	row := store.db.QueryRow("SELECT * FROM special_government_units WHERE id = $1", id)
	sgu, err := newSpecialGovernmentUnit(row)

	if err == nil {
		return sgu, nil
	}

	if err == sql.ErrNoRows {
		return sgu, fmt.Errorf("special government unit with id = %d not found: %w", id, err)
	}

	return sgu, fmt.Errorf("error executing query: %w", err)
}

func (store SpecialGovernmentUnit) FindByCode(code string) (models.SpecialGovernmentUnit, error) {
	row := store.db.QueryRow("SELECT * FROM special_government_units WHERE code = $1", code)
	sgu, err := newSpecialGovernmentUnit(row)

	if err == nil {
		return sgu, nil
	}

	if err == sql.ErrNoRows {
		return sgu, fmt.Errorf("special government unit with code = %d not found: %w", code, err)
	}

	return sgu, fmt.Errorf("error executing query: %w", err)
}

func (store SpecialGovernmentUnit) FindByName(name string) (models.SpecialGovernmentUnit, error) {
	row := store.db.QueryRow("SELECT * FROM special_government_units WHERE name = $1", name)
	sgu, err := newSpecialGovernmentUnit(row)

	if err == nil {
		return sgu, nil
	}

	if err == sql.ErrNoRows {
		return sgu, fmt.Errorf("special government unit with name = %d not found: %w", name, err)
	}

	return sgu, fmt.Errorf("error executing query: %w", err)
}

func (store SpecialGovernmentUnit) List(opts stores.SearchOpts) (stores.Collection[models.SpecialGovernmentUnit], error) {
	collection := stores.Collection[models.SpecialGovernmentUnit]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from special_government_units").Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM special_government_units ORDER BY $1 LIMIT $2 OFFSET $3",
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	sgus, err := newSpecialGovernmentUnits(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = sgus
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store SpecialGovernmentUnit) ListByProvinceCode(code string, opts stores.SearchOpts) (stores.Collection[models.SpecialGovernmentUnit], error) {
	collection := stores.Collection[models.SpecialGovernmentUnit]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from special_government_units WHERE province_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM special_government_units WHERE province_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	sgus, err := newSpecialGovernmentUnits(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = sgus
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newSpecialGovernmentUnits(rows *sql.Rows, count int) ([]models.SpecialGovernmentUnit, error) {
	sgus := make([]models.SpecialGovernmentUnit, 0, count)

	for rows.Next() {
		var s models.SpecialGovernmentUnit

		if err := rows.Scan(&s.Id, &s.Code, &s.Name, &s.ProvinceCode); err != nil {
			return sgus, err
		}
		sgus = append(sgus, s)
	}

	return sgus, nil
}

func newSpecialGovernmentUnit(row *sql.Row) (models.SpecialGovernmentUnit, error) {
	var sgu models.SpecialGovernmentUnit

	err := row.Scan(&sgu.Id, &sgu.Code, &sgu.Name, &sgu.ProvinceCode)

	return sgu, err
}
