package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
)

type MunicipalityStore struct {
	connection *sql.DB
}

func (store MunicipalityStore) Save(ctx context.Context, municipality models.Municipality) error {
	stmt, err := store.connection.PrepareContext(
		ctx,
		"INSERT INTO municipalities (code, name, income_class, population, province_id, district_id) VALUES ($1, $2, $3, $4, $5, $6)",
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
		municipality.ProvinceId,
		municipality.DistrictId,
	); err != nil {
		return fmt.Errorf("error executing query: %s", err)
	}

	return nil
}

func (store MunicipalityStore) Find(id int) (models.Municipality, error) {
	row := store.connection.QueryRow("SELECT * FROM municipalities WHERE municipalities.id = $1", id)
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
	row := store.connection.QueryRow("SELECT * FROM municipalities WHERE code = $1", code)
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
	row := store.connection.QueryRow("SELECT * FROM municipalities WHERE name = $1", name)
	municipality, err := newMunicipality(row)

	if err == nil {
		return municipality, nil
	}

	if err == sql.ErrNoRows {
		return municipality, fmt.Errorf("municipality with name = %s not found: %s", name, err)
	}

	return municipality, fmt.Errorf("error executing query: %s", err)
}

func newMunicipality(row *sql.Row) (models.Municipality, error) {
	var municipality models.Municipality

	err := row.Scan(
		&municipality.Id,
		&municipality.Code,
		&municipality.Name,
		&municipality.IncomeClass,
		&municipality.Population,
		&municipality.ProvinceId,
		&municipality.DistrictId,
	)

	if err != nil {
		return municipality, err
	}

	return municipality, nil
}
