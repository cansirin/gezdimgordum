package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/graph/model"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) AuthenticateUser(ctx context.Context, username string, password string) bool {
	user := model.User{}

	query := b.DB.Where("username = ?", username).First(&user)
	if query.Error != nil {
		return false
	}

	return models.CheckPasswordHash(password, user.Password)
}
