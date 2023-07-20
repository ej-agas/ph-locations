package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type SubMunicipality interface {
	Save(ctx context.Context, subMunicipality models.SubMunicipality) error
	Find(id int) (models.SubMunicipality, error)
	FindByCode(code string) (models.SubMunicipality, error)
	FindByName(name string) (models.SubMunicipality, error)
}
