package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) CreateUser(ctx context.Context, username string, password string) (*models.User, error) {
	hashedPassword, err := models.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Password: hashedPassword,
	}

	query := b.DB.Create(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	return &user, nil

}
