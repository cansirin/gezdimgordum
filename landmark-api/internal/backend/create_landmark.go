package backend

import (
	"context"
	"github.com/gezdimgordum/landmark-api/internal/models"
	"github.com/gosimple/slug"
)

func (b *PostgreSQLBackend) CreateLandmark(ctx context.Context, name string, description string, address string, stateID string, userID string) (*models.Landmark, error) {
	landmark := models.Landmark{
		Name:        name,
		Slug:        slug.MakeLang(name, "tr"),
		Description: description,
		Address:     address,
		StateID:     stateID,
		UserID:      userID,
	}

	query := b.DB.Create(&landmark)
	if query.Error != nil {
		return nil, query.Error
	}

	return &landmark, nil
}
