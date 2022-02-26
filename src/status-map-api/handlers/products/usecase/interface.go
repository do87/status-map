package usecase

import (
	"context"

	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
)

// Repository is the allowed usecase repo
type Repository interface {
	Reader
	Writer
}

// Reader represents reading functionality
type Reader interface {
	List(ctx context.Context) ([]entity.Product, error)
}

// Writer represents writing functionality
type Writer interface {
	Create(ctx context.Context, product entity.Product) (entity.Product, error)
	Update(ctx context.Context, product entity.Product) (entity.Product, error)
}
