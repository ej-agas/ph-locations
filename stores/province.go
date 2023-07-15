package stores

import (
	"context"
	"ph-locations/models"
)

type ProvinceStore interface {
	Save(ctx context.Context, province models.Province) error
	Find(id int) (models.Province, error)
	FindByCode(ctx context.Context, code string) (models.Province, error)
	FindByName(ctx context.Context, name string) (models.Province, error)
}
