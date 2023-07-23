package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
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
	var region models.Region

	row := store.db.QueryRow("SELECT * FROM regions WHERE id = $1", id)
	err := row.Scan(&region.Id, &region.Code, &region.Name, &region.Population)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with id = %d not found: %w", id, err)
	}

	return region, fmt.Errorf("error executing query: %w", err)
}

func (store RegionStore) FindByCode(code string) (models.Region, error) {
	var region models.Region

	row := store.db.QueryRow("SELECT * FROM regions WHERE code = $1", code)
	err := row.Scan(&region.Id, &region.Code, &region.Name, &region.Population)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with code = %s not found: %s", code, err)
	}

	return region, fmt.Errorf("error executing query: %s", err)
}

func (store RegionStore) FindByName(name string) (models.Region, error) {
	var region models.Region

	row := store.db.QueryRow("SELECT * FROM regions WHERE name = $1", name)
	err := row.Scan(&region.Id, &region.Code, &region.Name, &region.Population)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with name = %s not found: %s", name, err)
	}

	return region, fmt.Errorf("error executing query: %s", err)
}

func (store RegionStore) All() ([]models.Region, error) {
	regions := make([]models.Region, 0)

	rows, err := store.db.Query("SELECT * FROM regions")

	if err == sql.ErrNoRows {
		return regions, nil
	}

	defer rows.Close()

	for rows.Next() {
		region := models.Region{}

		if err := rows.Scan(&region.Id, &region.Code, &region.Name, &region.Population); err != nil {
			return regions, fmt.Errorf("error scanning row: %s", err)
		}

		regions = append(regions, region)
	}

	return regions, nil
}
