package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type BarangayStore interface {
	Save(ctx context.Context, barangay models.Barangay) error
	Find(id int) (models.Barangay, error)
	FindByCode(code string) (models.Barangay, error)
	FindByName(name string) (models.Barangay, error)
	ListByCityCode(code string, opts SearchOpts) (Collection[models.Barangay], error)
	ListByMunicipalityCode(code string, opts SearchOpts) (Collection[models.Barangay], error)
	ListBySubMunicipalityCode(code string, opts SearchOpts) (Collection[models.Barangay], error)
	ListBySpecialGovernmentUnitCode(code string, opts SearchOpts) (Collection[models.Barangay], error)
}
