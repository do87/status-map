package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/products/repository"
)

func DiscoverStages(ctx context.Context, cfg *config.Config, repo *repository.Repository) {
	fmt.Println("Discovering stages..")
	products, err := repo.List(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, product := range products {
		stages, err := getStages(ctx, cfg, product.Org, product.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		product = updateProductWithStages(product, stages)
		if _, err := repo.Update(ctx, product); err != nil {
			fmt.Println(err.Error())
			return
		}

		if i%5 == 0 {
			time.Sleep(time.Second * 2)
		}
	}

	fmt.Println("Job finished!")
}
