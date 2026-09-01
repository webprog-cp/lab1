package main

import (
	"cp_lab1/internal/config"
	"cp_lab1/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.New(cfg)

	host := cfg.Host + ":" + cfg.Port

	r.Run(host)
}
