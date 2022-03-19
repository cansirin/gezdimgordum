package backend

import (
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func (b *PostgreSQLBackend) GetUserID(username string) (string, error) {
	user := models.User{}

	query := b.DB.Where("username = ?", username).First(&user)
	if query.Error != nil {
		return "", query.Error
	}

	return user.ID.String(), nil
}
