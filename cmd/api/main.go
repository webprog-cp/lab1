package main

import (
	"log"

	"cp_lab1/internal/app"
	"cp_lab1/internal/config"
	"cp_lab1/internal/router"
)

// @title Cats shelter management API
// @version 1.0
// @description REST API service for cats shelter management.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.Load()
	a := app.New()

	r := router.New(cfg, a)

	host := cfg.Host + ":" + cfg.Port

	if err := r.Run(host); err != nil {
		log.Fatal(err)
	}
}
