package backend

import (
	"context"
	"github.com/gezdimgordum/landmark-api/internal/models"
	"github.com/gosimple/slug"
)

func (b *PostgreSQLBackend) CreateState(ctx context.Context, name string) (*models.State, error) {
	state := models.State{
		Name: name,
		Slug: slug.MakeLang(name, "tr"),
	}

	result := b.DB.Create(&state)
	if result.Error != nil {
		return nil, result.Error
	}

	return &state, nil
}
