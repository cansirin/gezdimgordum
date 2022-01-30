package backend

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	CreateLandmark(ctx context.Context, name string, description string, address string, stateID string, userID string) (*models.Landmark, error)
	GetLandmark(ctx context.Context, id string) (*models.Landmark, error)
	DeleteLandmark(ctx context.Context, id string) error
	UpdateLandmark(ctx context.Context, id string, name *string, description *string, address *string, state *string) error
	GetLandmarksByState(ctx context.Context, state string) ([]*models.Landmark, error)
	GetAllLandmarks(ctx context.Context) ([]*models.Landmark, error)
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{DB: db}
}
