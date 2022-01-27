package backend

import (
	"context"
	"github.com/gezdimgordum/landmark-api/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	CreateLandmark(ctx context.Context, name string, description string, address string, stateID string, userID string) (*models.Landmark, error)
	GetLandmark(ctx context.Context, id string) (*models.Landmark, error)
	DeleteLandmark(ctx context.Context, id string) error
	UpdateLandmark(ctx context.Context, id string, name *string, description *string, address *string, stateID *string) error
	GetLandmarksByStateID(ctx context.Context, id string) ([]*models.Landmark, error)

	CreateState(ctx context.Context, name string) (*models.State, error)
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{DB: db}
}
