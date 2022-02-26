package repository

import (
	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
	"github.com/do87/status-map/src/status-map-shared/pkg/database"
	"gorm.io/gorm"
)

// Repository service
type Repository struct {
	db *gorm.DB
}

// NewPostgres returns a new postgres service
func NewPostgres(db *database.Service) *Repository {
	if err := db.MigrateTables(&entity.Product{}); err != nil {
		panic(err)
	}

	return &Repository{
		db: db.GetDB(),
	}
}
