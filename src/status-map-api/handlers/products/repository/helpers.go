package repository

import "github.com/do87/status-map/src/status-map-api/handlers/products/entity"

func mergeStatusMaps(new, old entity.StatusMap) entity.StatusMap {
	merged := entity.StatusMap{}
	if old != nil {
		merged = old
	}

	for k, v := range new {
		merged[k] = v
	}
	return merged
}
