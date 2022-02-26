package jobs

import (
	"context"
	"fmt"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/products/repository"
)

func SyncProducts(ctx context.Context, cfg *config.Config, repo *repository.Repository, org string) {
	fmt.Printf("Starting product sync for %s\n", org)
	dbProducts, err := repo.List(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	products, err := getProducts(ctx, cfg, org)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	diff := getProductDiff(products.Items, dbProducts, org)
	for _, product := range diff {
		if _, err := repo.Create(ctx, product); err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Job finished!")
}
