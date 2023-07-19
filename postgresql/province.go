package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ej-agas/ph-locations/models"
)

type ProvinceStore struct {
	connection *sql.DB
}

func NewProvinceStore(connection *sql.DB) *ProvinceStore {
	return &ProvinceStore{connection: connection}
}

func (store ProvinceStore) Save(ctx context.Context, province models.Province) error {
	stmt, err := store.connection.PrepareContext(ctx, "INSERT INTO provinces (code, name, income_class, population, region_id) VALUES ($1, $2, $3)")

	if err != nil {
		return fmt.Errorf("error connecting to postgresql: %s", err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(province.Code, province.Name, province.IncomeClass, province.Population, province.RegionId); err != nil {
		return fmt.Errorf("error executing query: %s", err)
	}

	return nil
}

func (store ProvinceStore) Find(id int) (models.Province, error) {
	var province models.Province

	row := store.connection.QueryRow("SELECT * FROM provinces WHERE id = $1", id)
	err := row.Scan(&province.Id, &province.Code, &province.Name, &province.IncomeClass, &province.Population, &province.RegionId)

	if err == nil {
		return province, nil
	}

	if err == sql.ErrNoRows {
		return province, fmt.Errorf("province with id = %d not found: %s", id, err)
	}

	return province, fmt.Errorf("error executing query: %s", err)
}

func (store ProvinceStore) FindByCode(ctx context.Context, code string) (models.Province, error) {
	var province models.Province

	row := store.connection.QueryRow("SELECT * FROM provinces WHERE code = $1", code)
	err := row.Scan(&province.Id, &province.Code, &province.Name, &province.IncomeClass, &province.Population, &province.RegionId)

	if err == nil {
		return province, nil
	}

	if err == sql.ErrNoRows {
		return province, fmt.Errorf("province with code = %s not found: %s", code, err)
	}

	return province, fmt.Errorf("error executing query: %s", err)
}

func (store ProvinceStore) FindByName(ctx context.Context, name string) (models.Province, error) {
	var province models.Province

	row := store.connection.QueryRow("SELECT * FROM provinces WHERE name = $1", name)
	err := row.Scan(&province.Id, &province.Code, &province.Name, &province.IncomeClass, &province.Population, &province.RegionId)

	if err == nil {
		return province, nil
	}

	if err == sql.ErrNoRows {
		return province, fmt.Errorf("province with name = %s not found: %s", name, err)
	}

	return province, fmt.Errorf("error executing query: %s", err)
}
