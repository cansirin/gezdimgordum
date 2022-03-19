package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) GetUserLandmarks(ctx context.Context, userID string) ([]*models.Landmark, error) {
	var landmarks []*models.Landmark

	query := b.DB.Where("user_id = ?", userID).Find(&landmarks)
	if query.Error != nil {
		return nil, query.Error
	}

	return landmarks, nil

}
