package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/products/repository"
)

func DiscoverInfras(ctx context.Context, cfg *config.Config, repo *repository.Repository) {
	fmt.Println("Discovering infras..")
	products, err := repo.List(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, product := range products {
		infras, err := getInfras(ctx, cfg, product.Org, product.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		product = updateProductWithInfras(product, infras)
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
