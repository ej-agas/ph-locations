package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type DistrictStore interface {
	Save(ctx context.Context, district models.District) error
	Find(id int) (models.District, error)
	FindByCode(code string) (models.District, error)
	FindByName(name string) (models.District, error)
}
