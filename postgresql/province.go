package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
)

var (
	ErrInvalidCursor = errors.New("invalid cursor")
)

type ProvinceStore struct {
	db *sql.DB
}

func NewProvinceStore(connection *sql.DB) *ProvinceStore {
	return &ProvinceStore{db: connection}
}

func (store ProvinceStore) Save(ctx context.Context, province models.Province) error {
	stmt, err := store.db.PrepareContext(ctx, "INSERT INTO provinces (code, name, income_class, population, region_code) VALUES ($1, $2, $3, $4, $5)")

	if err != nil {
		return fmt.Errorf("error connecting to postgresql: %s", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(province.Code, province.Name, province.IncomeClass, province.Population, province.RegionCode); err != nil {
		return fmt.Errorf("error executing query: %s", err)
	}

	return nil
}

func (store ProvinceStore) Find(id int) (models.Province, error) {
	row := store.db.QueryRow("SELECT * FROM provinces WHERE id = $1", id)
	province, err := newProvince(row)

	if err == nil {
		return province, nil
	}

	if err == sql.ErrNoRows {
		return province, fmt.Errorf("province with id = %d not found: %s", id, err)
	}

	return province, fmt.Errorf("error executing query: %s", err)
}

func (store ProvinceStore) FindByCode(ctx context.Context, code string) (models.Province, error) {
	row := store.db.QueryRow("SELECT * FROM provinces WHERE code = $1", code)
	province, err := newProvince(row)

	if err == nil {
		return province, nil
	}

	if err == sql.ErrNoRows {
		return province, fmt.Errorf("province with code = %s not found: %s", code, err)
	}

	return province, fmt.Errorf("error executing query: %s", err)
}

func (store ProvinceStore) FindByName(ctx context.Context, name string) (models.Province, error) {
	row := store.db.QueryRow("SELECT * FROM provinces WHERE name = $1", name)
	province, err := newProvince(row)

	if err == nil {
		return province, nil
	}

	if err == sql.ErrNoRows {
		return province, fmt.Errorf("province with name = %s not found: %s", name, err)
	}

	return province, fmt.Errorf("error executing query: %s", err)
}

func (store ProvinceStore) List(opts stores.SearchOpts) (stores.Collection[models.Province], error) {
	collection := stores.Collection[models.Province]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from provinces").Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM provinces ORDER BY $1 LIMIT $2 OFFSET $3",
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	provinces, err := newProvinces(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = provinces
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func (store ProvinceStore) ListByRegionCode(code string, opts stores.SearchOpts) (stores.Collection[models.Province], error) {
	collection := stores.Collection[models.Province]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from provinces WHERE region_code = $1", code).Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	rows, err := store.db.Query(
		"SELECT * FROM provinces WHERE region_code = $1 ORDER BY $2 LIMIT $3 OFFSET $4",
		code,
		opts.Order,
		opts.Limit,
		offset,
	)

	if err != nil {
		return collection, err
	}

	provinces, err := newProvinces(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = provinces
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newProvinces(rows *sql.Rows, count int) ([]models.Province, error) {
	provinces := make([]models.Province, 0, count)

	for rows.Next() {
		var p models.Province

		if err := rows.Scan(
			&p.Id,
			&p.Code,
			&p.Name,
			&p.IncomeClass,
			&p.Population,
			&p.RegionCode,
		); err != nil {
			return provinces, err
		}
		provinces = append(provinces, p)
	}

	return provinces, nil
}

func newProvince(row *sql.Row) (models.Province, error) {
	var province models.Province

	err := row.Scan(
		&province.Id,
		&province.Code,
		&province.Name,
		&province.IncomeClass,
		&province.Population,
		&province.RegionCode,
	)

	return province, err
}
