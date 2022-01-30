package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) GetLandmark(ctx context.Context, id string) (*models.Landmark, error) {
	landmark := models.Landmark{}

	query := b.DB.First(&landmark, "id = ?", id)
	if query.Error != nil {
		return nil, query.Error
	}

	return &landmark, nil
}
