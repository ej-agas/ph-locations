package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"ph-locations/models"
)

type RegionStore struct {
	connection *sql.DB
}

func (store RegionStore) Save(ctx context.Context, region models.Region) error {
	stmt, err := store.connection.PrepareContext(ctx, "INSERT INTO regions (code, name, population) VALUES ($1, $2, $3)")

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

	row := store.connection.QueryRow("SELECT * FROM regions WHERE id = ?", id)
	err := row.Scan(&region.Id, &region.Code, &region.Name, &region.Population)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with id = %d not found: %s", id, err)
	}

	return region, fmt.Errorf("error executing query: %s", err)
}

func (store RegionStore) FindByCode(code string) (models.Region, error) {
	var region models.Region

	row := store.connection.QueryRow("SELECT * FROM regions WHERE code = ?", code)
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

	row := store.connection.QueryRow("SELECT * FROM regions WHERE name = ?", name)
	err := row.Scan(&region.Id, &region.Code, &region.Name, &region.Population)

	if err == nil {
		return region, nil
	}

	if err == sql.ErrNoRows {
		return region, fmt.Errorf("region with name = %s not found: %s", name, err)
	}

	return region, fmt.Errorf("error executing query: %s", err)
}
