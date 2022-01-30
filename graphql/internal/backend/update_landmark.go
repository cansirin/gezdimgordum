package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
	"github.com/gosimple/slug"
)

func (b *PostgreSQLBackend) UpdateLandmark(ctx context.Context, id string, name *string, description *string, address *string, state *string) error {
	landmark := models.Landmark{}
	if result := b.DB.First(&landmark, "id = ?", id); result.Error != nil {
		return result.Error
	}

	updates := models.Landmark{
		Name:        *name,
		Slug:        slug.MakeLang(*name, "tr"),
		Address:     *address,
		Description: *description,
		State:       *state,
	}

	if result := b.DB.Model(&landmark).Updates(&updates); result.Error != nil {
		return result.Error
	}

	return nil
}
