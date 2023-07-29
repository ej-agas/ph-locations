package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type CityStore interface {
	Save(ctx context.Context, city models.City) error
	Find(id int) (models.City, error)
	FindByCode(code string) (models.City, error)
	FindByName(name string) (models.City, error)
	List(opts SearchOpts) (Collection[models.City], error)
	ListByProvinceCode(code string, opts SearchOpts) (Collection[models.City], error)
	ListByDistrictCode(code string, opts SearchOpts) (Collection[models.City], error)
}
