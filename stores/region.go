package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type RegionStore interface {
	Save(ctx context.Context, region models.Region) error
	Find(id int) (models.Region, error)
	FindByCode(code string) (models.Region, error)
	FindByName(name string) (models.Region, error)
	List(opts SearchOpts) (Collection[models.Region], error)
}
