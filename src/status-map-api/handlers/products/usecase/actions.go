package usecase

import (
	"context"
	"net/http"

	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
)

// ListProducts returns a list of all products
func (s *Service) ListProducts(ctx context.Context, r *http.Request) (interface{}, error) {
	products, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	if isSetTo(r.URL.Query(), "groupBy", "platforms") {
		return groupProductsByPlatforms(products), nil
	}
	return products, nil
}

// CreateProduct creates a product
func (s *Service) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	return s.repo.Create(ctx, product)
}

// UpdateProduct updates a product
func (s *Service) UpdateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	return s.repo.Update(ctx, product)
}
