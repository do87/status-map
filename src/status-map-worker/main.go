package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/do87/status-map/src/status-map-api/config"
	"github.com/do87/status-map/src/status-map-api/handlers/health"
	"github.com/do87/status-map/src/status-map-api/handlers/products/repository"
	"github.com/do87/status-map/src/status-map-api/pkg/router"
	"github.com/do87/status-map/src/status-map-shared/pkg/database"
	"github.com/do87/status-map/src/status-map-worker/jobs"
	"github.com/go-co-op/gocron"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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
	repo := repository.NewPostgres(db)

	// serve default .well-known/live
	go serve(cfg)

	// define cron jobs
	ctx := context.Background()
	cron := gocron.NewScheduler(time.UTC)

	cron.Every(1).Day().SingletonMode().Do(jobs.SyncProducts, ctx, cfg, repo, "odjui")
	cron.Every(12).Hours().SingletonMode().Do(jobs.SyncProducts, ctx, cfg, repo, "sit")

	cron.Every(3).Hour().SingletonMode().Do(jobs.DiscoverStages, ctx, cfg, repo)
	cron.Every(4).Hour().SingletonMode().Do(jobs.DiscoverInfras, ctx, cfg, repo)

	cron.Every(7).Minute().SingletonMode().Do(jobs.GetInfraStatus, ctx, cfg, repo)
	cron.Every(5).Minute().SingletonMode().Do(jobs.GetStageStatus, ctx, cfg, repo)

	cron.StartBlocking()
}

func serve(cfg *config.Config) {
	r := router.New(cfg.IsLocal())
	health.Handle(r)

	// start server
	serverStr := fmt.Sprintf("%s:%d", cfg.Server.BindAddr, cfg.Server.HttpPort)
	fmt.Printf("Starting server at: http://%s\n", serverStr)
	if err := http.ListenAndServe(serverStr, r); err != nil {
		panic(err)
	}

	fmt.Println("Goodbye!")
}
