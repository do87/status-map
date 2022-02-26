package presenter

import (
	"sort"

	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
)

// MapEntityProduct maps entity product to presenter
func MapEntityProduct(p entity.Product) Product {
	product := Product{}
	product.Etag = generateEtag(p)
	product.ID = p.ID
	product.Org = p.Org
	product.Platform = p.Platform
	product.CreatedAt = p.CreatedAt
	product.UpdatedAt = p.UpdatedAt
	product.Infras = trimInfras(p.Infras)
	product.Stages = p.Stages
	product.StatusInfras = MapEntityStatusMap(trimStatusInfras(p.StatusInfra))
	product.StatusStages = MapEntityStatusMap(p.StatusStages)
	return product
}

// MapProductsEntity maps entity data to presenter types
func MapEntityProducts(data []entity.Product) Products {
	p := Products{}
	p.Kind = kindProducts
	p.Items = make([]Product, 0)

	for _, v := range data {
		item := MapEntityProduct(v)
		p.Items = append(p.Items, item)
	}

	return p
}

// MapEntityPlatforms maps platform entity data to presenter types
func MapEntityPlatforms(data map[string][]entity.Product) ProductsGroupedByPlatforms {
	p := ProductsGroupedByPlatforms{
		Kind:  kindProductsGroupedByPlatforms,
		Items: make([]Platform, 0),
	}
	platforms := make([]string, 0)
	for name := range data {
		platforms = append(platforms, name)
	}
	sort.Strings(platforms)
	for _, name := range platforms {
		item := MapEntityProductsToPlatform(name, data[name])
		p.Items = append(p.Items, item)
	}
	return p
}

// MapEntityProductsToPlatform maps products entity data, to a platform
func MapEntityProductsToPlatform(name string, data []entity.Product) Platform {
	p := Platform{
		Kind:  kindPlatform,
		Name:  name,
		Items: make([]Product, 0),
	}
	for _, v := range data {
		item := MapEntityProduct(v)
		p.Items = append(p.Items, item)
	}
	return p
}

// MapEntityStatusMap maps entity status map to presenter status map
func MapEntityStatusMap(entStatusMap entity.StatusMap) StatusMap {
	statusMap := StatusMap{}
	for name, s := range entStatusMap {
		statusMap[name] = Status{
			CurrentStatus: s.CurrentStatus,
			UpdatedAt:     s.UpdatedAt,
		}
	}
	return statusMap
}
