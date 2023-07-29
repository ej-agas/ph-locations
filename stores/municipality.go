package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type MunicipalityStore interface {
	Save(ctx context.Context, municipality models.Municipality) error
	Find(id int) (models.Municipality, error)
	FindByCode(code string) (models.Municipality, error)
	FindByName(name string) (models.Municipality, error)
	List(opts SearchOpts) (Collection[models.Municipality], error)
	ListByProvinceCode(code string, opts SearchOpts) (Collection[models.Municipality], error)
	ListByDistrictCode(code string, opts SearchOpts) (Collection[models.Municipality], error)
}
