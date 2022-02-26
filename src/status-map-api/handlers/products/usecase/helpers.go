package usecase

import (
	"net/url"

	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
)

// isSetTo checks the request if a param is set to a given value
func isSetTo(params url.Values, param, value string) bool {
	if groups, ok := params[param]; ok {
		for _, v := range groups {
			if v == value {
				return true
			}
		}
	}
	return false
}

// groupProductsByPlatforms returns a list of platforms and the products that belong to them
func groupProductsByPlatforms(products []entity.Product) map[string][]entity.Product {
	platforms := map[string][]entity.Product{}
	for _, product := range products {
		platform := product.Platform
		if _, ok := platforms[platform]; !ok {
			platforms[platform] = make([]entity.Product, 0)
		}
		platforms[platform] = append(platforms[platform], product)
	}
	return platforms
}
