package router

import (
	"cp_lab1/internal/app"
	"cp_lab1/internal/config"
	"cp_lab1/internal/handlers"

	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config, app *app.App) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	ge := gin.New()

	registerRoutes(ge)

	ge.Use(gin.Logger())
	ge.Use(gin.Recovery()) // Recovers from any panic and returns 500 if there is one

	return ge
}

func registerCats(ge *gin.Engine, ch *handlers.CatHandler) {
	cats := ge.Group("/cats")
	cats.GET("", ch.GetAll)
}

func registerShelters(ge *gin.Engine, sh *handlers.ShelterHandler) {
	//#TODO: finish app, split logic
}

func registerRoutes(ge *gin.Engine, app *app.App) {
	registerCats(ge, app.CatHandler)
	registerShelters(ge, app.ShelterHandler)
}
