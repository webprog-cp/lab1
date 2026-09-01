package router

import (
	"cp_lab1/internal/config"

	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *gin.Engine {
	ge := gin.New()

	gin.SetMode(cfg.GinMode)

	ge.Use(gin.Logger())
	ge.Use(gin.Recovery()) // Recovers from any panic and returns 500 if there is one

	return ge
}

func registerCats(ge *gin.Engine) *gin.Engine {
	
}

func registerShelters(ge *gin.Engine) *gin.Engine {
	
}

func registerRoutes(ge *gin.Engine) {
	registerCats(ge)
	registerShelters(ge)
}
