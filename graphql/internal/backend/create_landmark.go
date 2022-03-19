package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
	"github.com/gosimple/slug"
)

func (b *PostgreSQLBackend) CreateLandmark(ctx context.Context, name string, description string, address string, state string, userID string) (*models.Landmark, error) {
	landmark := models.Landmark{
		Name:        name,
		Slug:        slug.MakeLang(name, "tr"),
		Description: description,
		Address:     address,
		State:       state,
		UserID:      userID,
	}

	result := b.DB.Create(&landmark)
	if result.Error != nil {
		return nil, result.Error
	}

	return &landmark, nil
}
