package stores

import (
	"context"
	"ph-locations/models"
)

type RegionStore interface {
	Save(ctx context.Context, region models.Region) error
	Find(id int) (models.Region, error)
	FindByCode(ctx context.Context, code string) (models.Region, error)
	FindByName(ctx context.Context, name string) (models.Region, error)
}
