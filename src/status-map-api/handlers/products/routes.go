package products

import (
	"net/http"

	"github.com/do87/status-map/src/status-map-api/handlers/products/presenter"
	"github.com/do87/status-map/src/status-map-api/handlers/products/usecase"
	"github.com/do87/status-map/src/status-map-shared/pkg/common"
	"github.com/go-chi/render"
)

// setRoutes attaches product routes
func (p *Products) setRoutes() *Products {
	p.route.Get("/products", listProducts(p.service))
	return p
}

func listProducts(service *usecase.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := service.ListProducts(r.Context(), r)
		if err != nil {
			common.ErrorResponse(w, r, http.StatusInternalServerError, err)
		}
		render.JSON(w, r, presenter.PresentProductList(data))
	}
}
