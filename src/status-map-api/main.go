package main

import (
	"fmt"
	"log"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-shared/pkg/database"
)

func main() {
	fmt.Println("Loading configuration")
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	fmt.Println("Setting up database connection")
	db, err := database.New(cfg.Database, cfg.IsLocal())
	if err != nil {
		log.Fatalf("db init error: %v", err)
	}
	defer db.Close()

	if err := run(cfg, db); err != nil {
		panic(err)
	}
}
