package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) GetLandmarksByState(ctx context.Context, state string) ([]*models.Landmark, error) {
	var landmarks []*models.Landmark

	query := b.DB.Find(&landmarks, "state = ?", state)
	if query.Error != nil {
		return nil, query.Error
	}

	return landmarks, nil
}
