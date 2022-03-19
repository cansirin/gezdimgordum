package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	CreateLandmark(ctx context.Context, name string, description string, address string, stateID string, userID string) (*models.Landmark, error)
	GetLandmark(ctx context.Context, id string) (*models.Landmark, error)
	DeleteLandmark(ctx context.Context, id string, userID string) error
	UpdateLandmark(ctx context.Context, id string, name *string, description *string, address *string, state *string) error
	GetLandmarksByState(ctx context.Context, state string) ([]*models.Landmark, error)
	GetAllLandmarks(ctx context.Context) ([]*models.Landmark, error)

	CreateUser(ctx context.Context, username string, password string) (*models.User, error)
	GetUserID(username string) (string, error)
	AuthenticateUser(ctx context.Context, username string, password string) bool
	GetUser(ctx context.Context, id string) (*models.User, error)
	GetUserLandmarks(ctx context.Context, userID string) ([]*models.Landmark, error)
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{DB: db}
}
