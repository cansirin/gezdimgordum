package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) GetUser(ctx context.Context, id string) (*models.User, error) {
	user := models.User{}

	query := b.DB.First(&user, id)
	if query.Error != nil {
		return nil, query.Error
	}

	return &user, nil
}
