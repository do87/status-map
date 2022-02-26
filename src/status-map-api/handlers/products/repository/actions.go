package repository

import (
	"context"
	"errors"

	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
	"github.com/do87/status-map/src/status-map-shared/pkg/common"
)

// Get returns record by ID
func (r *Repository) Get(ctx context.Context, id, org string) (product entity.Product, err error) {
	result := r.db.First(&product, "id = ? AND org = ?", id, org)
	if err = result.Error; err != nil {
		return
	}
	if result.RowsAffected == 0 {
		err = errors.New(common.ERR_NOT_FOUND)
	}
	return
}

// List returns all products
func (r *Repository) List(ctx context.Context) (products []entity.Product, err error) {
	result := r.db.Order("id ASC").Find(&products)
	if result.Error != nil {
		return products, result.Error
	}
	return products, nil
}

// Create a product if not exists
func (r *Repository) Create(ctx context.Context, product entity.Product) (entity.Product, error) {
	result := r.db.FirstOrCreate(&product)
	if result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

// Update an existing product
func (r *Repository) Update(ctx context.Context, product entity.Product) (entity.Product, error) {
	existing, err := r.Get(ctx, product.ID, product.Org)
	if err != nil {
		return entity.Product{}, err
	}

	product.StatusInfra = mergeStatusMaps(product.StatusInfra, existing.StatusInfra)
	product.StatusStages = mergeStatusMaps(product.StatusStages, existing.StatusStages)
	if err := r.db.Model(&existing).Updates(&product).Error; err != nil {
		return existing, err
	}
	return product, nil
}
