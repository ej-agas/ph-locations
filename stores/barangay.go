package stores

import (
	"context"
	"github.com/ej-agas/ph-locations/models"
)

type Barangay interface {
	Save(ctx context.Context, barangay models.Barangay) error
	Find(id int) (models.Barangay, error)
	FindByCode(code string) (models.Barangay, error)
	FindByName(name string) (models.Barangay, error)
}
