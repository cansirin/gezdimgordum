package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type State struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name string
	Slug string `gorm:"index"`
}

type Landmark struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string
	Slug        string
	Address     string
	Description string

	StateID string
	State   State

	UserID string
}

func AutoMigrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&State{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Landmark{}); err != nil {
		return err
	}

	return nil
}