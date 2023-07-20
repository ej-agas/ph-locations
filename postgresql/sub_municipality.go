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
}
