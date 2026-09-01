package router

import (
	"cp_lab1/internal/config"
	"cp_lab1/internal/handlers"

	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	ge := gin.New()

	ge.Use(gin.Logger())
	ge.Use(gin.Recovery()) // Recovers from any panic and returns 500 if there is one

	return ge
}

func registerCats(ge *gin.Engine, ch *handlers.CatHandler) *gin.Engine {
	cats := ge.Group("/cats")
	cats.GET("", ch.GetAll)
	return ge
}

func registerShelters(ge *gin.Engine) *gin.Engine {

}

func registerRoutes(ge *gin.Engine) {
	registerCats(ge)
	registerShelters(ge)
}
