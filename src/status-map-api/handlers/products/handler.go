package products

import (
	"github.com/do87/status-map/src/status-map-api/handlers/products/repository"
	"github.com/do87/status-map/src/status-map-api/handlers/products/usecase"
	"github.com/do87/status-map/src/status-map-shared/pkg/database"
	"github.com/go-chi/chi/v5"
)

type Products struct {
	route   *chi.Mux
	repo    *repository.Repository
	service *usecase.Service
}

func Handle(r *chi.Mux, db *database.Service) {
	repo := repository.NewPostgres(db)
	p := &Products{
		route:   r,
		repo:    repo,
		service: usecase.NewService(repo),
	}

	p.setRoutes()
}
