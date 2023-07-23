package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
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
		"INSERT INTO special_government_units (code, name, province_id) VALUES ($1, $2, $3)",
	)

	if err != nil {
		return fmt.Errorf("error creating prepared statement: %w", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(sgu.Code, sgu.Name, sgu.ProvinceId); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (store SpecialGovernmentUnit) Find(id int) (models.SpecialGovernmentUnit, error) {
	var sgu models.SpecialGovernmentUnit

	row := store.db.QueryRow("SELECT * FROM special_government_units WHERE id = $1", id)
	err := row.Scan(&sgu.Id, &sgu.Code, &sgu.Name, &sgu.ProvinceId)

	if err == nil {
		return sgu, nil
	}

	if err == sql.ErrNoRows {
		return sgu, fmt.Errorf("special government unit with id = %d not found: %w", id, err)
	}

	return sgu, fmt.Errorf("error executing query: %w", err)
}

func (store SpecialGovernmentUnit) FindByCode(code string) (models.SpecialGovernmentUnit, error) {
	var sgu models.SpecialGovernmentUnit

	row := store.db.QueryRow("SELECT * FROM special_government_units WHERE code = $1", code)
	err := row.Scan(&sgu.Id, &sgu.Code, &sgu.Name, &sgu.ProvinceId)

	if err == nil {
		return sgu, nil
	}

	if err == sql.ErrNoRows {
		return sgu, fmt.Errorf("special government unit with code = %d not found: %w", code, err)
	}

	return sgu, fmt.Errorf("error executing query: %w", err)
}

func (store SpecialGovernmentUnit) FindByName(name string) (models.SpecialGovernmentUnit, error) {
	var sgu models.SpecialGovernmentUnit

	row := store.db.QueryRow("SELECT * FROM special_government_units WHERE name = $1", name)
	err := row.Scan(&sgu.Id, &sgu.Code, &sgu.Name, &sgu.ProvinceId)

	if err == nil {
		return sgu, nil
	}

	if err == sql.ErrNoRows {
		return sgu, fmt.Errorf("special government unit with name = %d not found: %w", name, err)
	}

	return sgu, fmt.Errorf("error executing query: %w", err)
}
