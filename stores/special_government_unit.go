package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type SpecialGovernmentUnit interface {
	Save(ctx context.Context, subMunicipality models.SpecialGovernmentUnit) error
	Find(id int) (models.SpecialGovernmentUnit, error)
	FindByCode(code string) (models.SpecialGovernmentUnit, error)
	FindByName(name string) (models.SpecialGovernmentUnit, error)
	ListByProvinceCode(code string, opts SearchOpts) (Collection[models.SpecialGovernmentUnit], error)
}
