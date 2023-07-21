package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
)

type SubMunicipalityStore struct {
	db *sql.DB
}

func (store SubMunicipalityStore) Save(ctx context.Context, subMunicipality models.SubMunicipality) error {
	stmt, err := store.db.PrepareContext(
		ctx,
		"INSERT INTO sub_municipalities (code, name, population, city_id) VALUES ($1, $2, $3, $4)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %w", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(
		subMunicipality.Code,
		subMunicipality.Name,
		subMunicipality.Population,
		subMunicipality.CityId,
	); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (store SubMunicipalityStore) Find(id int) (models.SubMunicipality, error) {
	var subMunicipality models.SubMunicipality

	row := store.db.QueryRow("SELECT * FROM sub_municipalities WHERE id = $1", id)
	err := row.Scan(&subMunicipality.Id, &subMunicipality.Code, &subMunicipality.Name, &subMunicipality.Population)

	if err == nil {
		return subMunicipality, nil
	}

	if err == sql.ErrNoRows {
		return subMunicipality, fmt.Errorf("SubMunicipality with id = %d not found: %w", id, err)
	}

	return subMunicipality, fmt.Errorf("error executing query: %w", err)
}

func (store SubMunicipalityStore) FindByCode(code string) (models.SubMunicipality, error) {
	var subMunicipality models.SubMunicipality

	row := store.db.QueryRow("SELECT * FROM sub_municipalities WHERE code = $1", code)
	err := row.Scan(&subMunicipality.Id, &subMunicipality.Code, &subMunicipality.Name, &subMunicipality.Population)

	if err == nil {
		return subMunicipality, nil
	}

	if err == sql.ErrNoRows {
		return subMunicipality, fmt.Errorf("SubMunicipality with code = %s not found: %w", code, err)
	}

	return subMunicipality, fmt.Errorf("error executing query: %w", err)
}

func (store SubMunicipalityStore) FindByName(name string) (models.SubMunicipality, error) {
	var subMunicipality models.SubMunicipality

	row := store.db.QueryRow("SELECT * FROM sub_municipalities WHERE name = $1", name)
	err := row.Scan(&subMunicipality.Id, &subMunicipality.Code, &subMunicipality.Name, &subMunicipality.Population)

	if err == nil {
		return subMunicipality, nil
	}

	if err == sql.ErrNoRows {
		return subMunicipality, fmt.Errorf("SubMunicipality with name = %s not found: %w", name, err)
	}

	return subMunicipality, fmt.Errorf("error executing query: %w", err)
}
