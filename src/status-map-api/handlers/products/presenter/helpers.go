package presenter

import (
	"fmt"
	"strings"

	"github.com/amalfra/etag"
	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
)

const (
	// kinds
	kindPlatform                   = "labs:statusmap:platforms"
	kindProducts                   = "labs:statusmap:products"
	kindProductsGroupedByPlatforms = "labs:statusmap:products;group-by=platforms"
)

// PresentProductList handles presenting a list of products
// and cases where the products are grouped
func PresentProductList(data interface{}) interface{} {
	if v, ok := data.([]entity.Product); ok {
		return MapEntityProducts(v)
	}

	if v, ok := data.(map[string][]entity.Product); ok {
		return MapEntityPlatforms(v)
	}
	return nil
}

func generateEtag(data entity.Product) string {
	d := fmt.Sprintf("%s %s %s", data.ID, data.Org, data.UpdatedAt)
	tag := etag.Generate(d, true)
	s := strings.Split(tag, `"`)
	if len(s) > 2 {
		return s[1]
	}
	return tag
}

func trimInfras(infras []string) []string {
	t := []string{}
	for _, i := range infras {
		if len(i) >= 4 {
			i = i[0:4]
		}
		t = append(t, i)
	}
	return t
}

func trimStatusInfras(sis entity.StatusMap) entity.StatusMap {
	newMap := entity.StatusMap{}
	for k, v := range sis {
		if len(k) > 4 {
			k = k[0:4]
		}
		newMap[k] = v
	}
	return newMap
}
