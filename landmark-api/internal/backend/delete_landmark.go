package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/landmark-api/internal/models"
)

func (b *PostgreSQLBackend) DeleteLandmark(ctx context.Context, id string) error {
	landmark := models.Landmark{}
	result := b.DB.First(&landmark, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = b.DB.Delete(&landmark)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
