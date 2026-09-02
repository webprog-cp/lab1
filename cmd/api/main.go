package main

import (
	"cp_lab1/internal/app"
	"cp_lab1/internal/config"
	"cp_lab1/internal/router"
)

func main() {
	cfg := config.Load()
	a := app.New()

	r := router.New(cfg, a)

	host := cfg.Host + ":" + cfg.Port

	r.Run(host)
}
