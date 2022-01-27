package backend

import (
	"context"
	"github.com/gezdimgordum/landmark-api/internal/models"
)

func (b *PostgreSQLBackend) GetLandmark(ctx context.Context, id string) (*models.Landmark, error) {
	landmark := models.Landmark{}

	query := b.DB.First(&landmark, "id = ?", id)
	if query.Error != nil {
		return nil, query.Error
	}

	query = b.DB.Preload("State").Find(&landmark)

	return &landmark, nil
}
