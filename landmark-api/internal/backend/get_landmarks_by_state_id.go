package backend

import (
	"context"
	"github.com/gezdimgordum/landmark-api/internal/models"
)

func (b *PostgreSQLBackend) GetLandmarksByStateID(ctx context.Context, id string) ([]*models.Landmark, error) {
	var landmarks []*models.Landmark

	query := b.DB.Preload("State").Find(&landmarks, "state_id = ?", id)
	if query.Error != nil {
		return nil, query.Error
	}

	return landmarks, nil
}
