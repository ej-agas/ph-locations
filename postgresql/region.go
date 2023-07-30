package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
	"github.com/ej-agas/ph-locations/stores"
	"math"
)

type RegionStore struct {
	db *sql.DB
}

func NewRegionStore(connection *sql.DB) *RegionStore {
	return &RegionStore{db: connection}
}

func (store RegionStore) Save(ctx context.Context, region models.Region) error {
	stmt, err := store.db.PrepareContext(ctx, "INSERT INTO regions (code, name, population) VALUES ($1, $2, $3)")

	if err != nil {
		return fmt.Errorf("error connecting to postgresql: %s", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(region.Code, region.Name, region.Population); err != nil {
		return fmt.Errorf("error executing query: %s", err)
	}

	return nil
}

func (store RegionStore) Find(id int) (models.Region, error) {
	row := store.db.QueryRow("SELECT * FROM regions WHERE id = $1", id)
	region, err := newRegion(row)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with id = %d not found: %w", id, err)
	}

	return region, fmt.Errorf("error executing query: %w", err)
}

func (store RegionStore) FindByCode(code string) (models.Region, error) {
	row := store.db.QueryRow("SELECT * FROM regions WHERE code = $1", code)
	region, err := newRegion(row)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with code = %s not found: %s", code, err)
	}

	return region, fmt.Errorf("error executing query: %s", err)
}

func (store RegionStore) FindByName(name string) (models.Region, error) {
	row := store.db.QueryRow("SELECT * FROM regions WHERE name = $1", name)
	region, err := newRegion(row)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with name = %s not found: %s", name, err)
	}

	return region, fmt.Errorf("error executing query: %s", err)
}

func (store RegionStore) List(opts stores.SearchOpts) (stores.Collection[models.Region], error) {
	collection := stores.Collection[models.Region]{}
	var totalRows int64
	var totalPages float64
	offset := (opts.Page - 1) * opts.Limit

	err := store.db.QueryRow("SELECT count(id) from regions").Scan(&totalRows)
	if err != nil {
		return collection, err
	}

	totalPages = math.Ceil(float64(totalRows) / float64(opts.Limit))
	if totalRows < int64(opts.Limit) {
		totalPages = 1
	}

	q := fmt.Sprintf("SELECT * FROM regions ORDER BY %s %s LIMIT $1 OFFSET $2", opts.Order, opts.Sort)
	rows, err := store.db.Query(q, opts.Limit, offset)

	if err != nil {
		return collection, err
	}

	regions, err := newRegions(rows, opts.Limit)
	if err != nil {
		return collection, err
	}

	paginationInfo := stores.PaginationInfo{
		TotalPages:  int(totalPages),
		PerPage:     opts.Limit,
		CurrentPage: opts.Page,
	}

	collection.Data = regions
	collection.PaginationInfo = paginationInfo

	return collection, nil
}

func newRegions(rows *sql.Rows, count int) ([]models.Region, error) {
	regions := make([]models.Region, 0, count)

	for rows.Next() {
		var r models.Region

		if err := rows.Scan(&r.Id, &r.Code, &r.Name, &r.Population); err != nil {
			return regions, err
		}
		fmt.Printf("%#v\n", r)
		regions = append(regions, r)
	}

	return regions, nil
}

func newRegion(row *sql.Row) (models.Region, error) {
	var region models.Region

	err := row.Scan(&region.Id, &region.Code, &region.Name, &region.Population)

	return region, err
}
