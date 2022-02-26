package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/products/entity"
	"github.com/do87/status-map/src/status-map-shared/pkg/iapclient"
	"github.com/do87/status-map/src/status-map-shared/pkg/payload"
)

const (
	apiProducts    = "https://odj.cloud/api/orgs/%s/products"
	apiInfras      = "https://odj.cloud/api/orgs/%s/products/%s/infras"
	apiInfraStatus = "https://odj.cloud/api/orgs/%s/products/%s/infras/%s/status"
	apiStages      = "https://odj.cloud/api/orgs/%s/products/%s/stages"
	apiStageStatus = "https://odj.cloud/api/orgs/%s/products/%s/stages/%s/status"
)

func getProducts(ctx context.Context, cfg *config.Config, org string) (res payload.Products, err error) {
	fmt.Printf("- fetching products for org%s\n", org)
	api := fmt.Sprintf(apiProducts, org)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		return payload.Products{}, err
	}

	client := iapclient.New(cfg.IAP.ClinetID)
	out, err := client.Apply(ctx, req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(out, &res); err != nil {
		fmt.Printf("unmarshal failed %v", err)
		return
	}

	fmt.Printf("- found %d products\n", len(res.Items))
	return res, nil
}

func getInfras(ctx context.Context, cfg *config.Config, org, productID string) (res payload.InfraList, err error) {
	fmt.Printf("- fetching infra for %s:%s", org, productID)
	api := fmt.Sprintf(apiInfras, org, productID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		return payload.InfraList{}, err
	}
	client := iapclient.New(cfg.IAP.ClinetID)
	out, err := client.Apply(ctx, req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(out, &res); err != nil {
		fmt.Printf(" ... unmarshal failed: %v\n", err)
		return
	}
	fmt.Printf(" ... done.\n")
	return res, nil
}

func getStages(ctx context.Context, cfg *config.Config, org, productID string) (res payload.StageList, err error) {
	fmt.Printf("- fetching stages for %s:%s", org, productID)
	api := fmt.Sprintf(apiStages, org, productID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		return payload.StageList{}, err
	}
	client := iapclient.New(cfg.IAP.ClinetID)
	out, err := client.Apply(ctx, req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(out, &res); err != nil {
		fmt.Printf(" ... unmarshal failed: %v\n", err)
		return
	}
	fmt.Printf(" ... done.\n")
	return res, nil
}

func getInfraStatus(ctx context.Context, cfg *config.Config, org, product, infra string) (status string, err error) {
	fmt.Printf("getting infra status for %s:%s:%s", org, product, infra)
	api := fmt.Sprintf(apiInfraStatus, org, product, infra)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		fmt.Printf(" ... failed: %v\n", err)
		return
	}
	client := iapclient.New(cfg.IAP.ClinetID)
	out, err := client.Apply(ctx, req)
	status = string(out)
	if !isValidStatus(status) {
		status = "unknown"
	}
	fmt.Printf(" ... done.\n")
	return
}

func getStageStatus(ctx context.Context, cfg *config.Config, org, product, stage string) (status string, err error) {
	fmt.Printf("getting stage status for %s:%s:%s", org, product, stage)
	api := fmt.Sprintf(apiStageStatus, org, product, stage)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		fmt.Printf(" ... failed: %v\n", err)
		return
	}
	client := iapclient.New(cfg.IAP.ClinetID)
	out, err := client.Apply(ctx, req)
	status = string(out)
	if !isValidStatus(status) {
		status = "unknown"
	}
	fmt.Printf(" ... done.\n")
	return
}

func isValidStatus(status string) bool {
	for _, a := range []string{"done", "error", "unknown", "canceled", "running"} {
		if a == status {
			return true
		}
	}
	return false
}

func getProductDiff(a []payload.Product, b []entity.Product, org string) map[string]entity.Product {
	diff := map[string]entity.Product{}
	for _, product := range a {
		diff[product.Name] = entity.Product{
			ID:  product.Name,
			Org: org,
		}
	}
	for _, product := range b {
		delete(diff, product.ID)
	}
	return diff
}

// updateProductWithInfras adds infras to the product and the platform name
func updateProductWithInfras(product entity.Product, infras payload.InfraList) entity.Product {
	all := []string{}
	platform := "unknown"
	for _, infra := range infras {
		all = append(all, infra.Environment)
		test := ""
		trimmedName := infra.Environment
		if len(infra.Environment) >= 4 {
			test = infra.Environment[0:2]
			trimmedName = infra.Environment[0:4]
		}
		statusInfras, ok := product.StatusInfra[infra.Environment]
		if ok {
			delete(product.StatusInfra, infra.Environment)
			product.StatusInfra[trimmedName] = statusInfras
		}
		switch test {
		case "az":
			if platform == "unknown" || platform == "azure" {
				platform = "azure"
			} else {
				platform = "mixed"
			}
		case "si":
			if platform == "unknown" || platform == "stackit" {
				platform = "stackit"
			} else {
				platform = "mixed"
			}
		case "li":
			fallthrough
		case "lv":
			fallthrough
		case "wo":
			fallthrough
		case "wk":
			if platform == "unknown" || platform == "gcp" {
				platform = "gcp"
			} else {
				platform = "mixed"
			}
		}

	}
	if platform == "mixed" {
		platform = "unknown"
	}

	product.Platform = platform
	product.Infras = all
	return product
}

// updateProductWithStages adds stages to the product
func updateProductWithStages(product entity.Product, stages payload.StageList) entity.Product {
	names := []string{}
	for _, stage := range stages {
		names = append(names, stage.Name)
	}
	product.Stages = names
	return product
}
