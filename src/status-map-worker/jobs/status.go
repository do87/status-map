package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
	"github.com/do87/status-map/src/status-map-api/handlers/products/repository"
)

func GetInfraStatus(ctx context.Context, cfg *config.Config, repo *repository.Repository) {
	fmt.Println("Get infra status for every product")
	productList, err := repo.List(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, product := range productList {
		statusMap := entity.StatusMap{}
		for k, v := range product.StatusInfra {
			statusMap[k] = entity.Status{
				CurrentStatus: v.CurrentStatus,
				UpdatedAt:     v.UpdatedAt,
			}
		}

		for _, infra := range product.Infras {
			status, err := getInfraStatus(ctx, cfg, product.Org, product.ID, infra)
			if err != nil {
				fmt.Println(err)
				continue
			}
			statusMap[infra] = entity.Status{
				CurrentStatus: status,
				UpdatedAt:     time.Now(),
			}
		}

		product.StatusInfra = statusMap
		if _, err := repo.Update(ctx, product); err != nil {
			fmt.Println(err.Error())
		}

		if i%20 == 0 {
			time.Sleep(time.Second * 2)
		}
	}

	fmt.Println("Job finished!")
}

func GetStageStatus(ctx context.Context, cfg *config.Config, repo *repository.Repository) {
	fmt.Println("Get stage status for every product")
	productList, err := repo.List(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, product := range productList {
		statusMap := entity.StatusMap{}
		for k, v := range product.StatusStages {
			statusMap[k] = entity.Status{
				CurrentStatus: v.CurrentStatus,
				UpdatedAt:     v.UpdatedAt,
			}
		}

		for _, stage := range product.Stages {
			status, err := getStageStatus(ctx, cfg, product.Org, product.ID, stage)
			if err != nil {
				fmt.Println(err)
				continue
			}
			statusMap[stage] = entity.Status{
				CurrentStatus: status,
				UpdatedAt:     time.Now(),
			}
		}

		product.StatusStages = statusMap
		if _, err := repo.Update(ctx, product); err != nil {
			fmt.Println(err.Error())
		}

		if i%20 == 0 {
			time.Sleep(time.Second * 2)
		}
	}

	fmt.Println("Job finished!")
}
