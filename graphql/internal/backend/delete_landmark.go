package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) DeleteLandmark(ctx context.Context, id string, userID string) error {
	user := models.User{}

	query := b.DB.First(&user, userID)
	if query.Error != nil {
		return query.Error
	}

	err := b.DB.Model(&user).Association("Landmarks").Delete(user)
	if err != nil {
		return err
	}

	return nil
}
