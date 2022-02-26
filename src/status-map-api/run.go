package main

import (
	"fmt"
	"net/http"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/health"
	"github.com/do87/status-map/src/status-map-api/handlers/products"
	"github.com/do87/status-map/src/status-map-api/pkg/router"
	"github.com/do87/status-map/src/status-map-shared/pkg/database"
)

// run applies the handlers and starts the API server
func run(cfg *config.Config, db *database.Service) error {
	fmt.Println("Creating router")
	r := router.New(cfg.IsLocal())

	fmt.Println("Running handlers for components")
	health.Handle(r)
	products.Handle(r, db)

	// start server
	serverStr := fmt.Sprintf("%s:%d", cfg.Server.BindAddr, cfg.Server.HttpPort)
	fmt.Printf("Starting server at: http://%s\n", serverStr)
	return http.ListenAndServe(serverStr, r)
}
