package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/landmark-api/internal/models"
)

func (b *PostgreSQLBackend) GetAllLandmarks(ctx context.Context) ([]*models.Landmark, error) {
	var landmarks []*models.Landmark

	result := b.DB.Find(&landmarks)
	if result.Error != nil {
		return nil, result.Error
	}

	return landmarks, nil
}
