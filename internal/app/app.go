package app

import (
	h "cp_lab1/internal/handlers"
	r "cp_lab1/internal/repositories"
)

type App struct {
	CatHandler     *h.CatHandler
	ShelterHandler *h.ShelterHandler
}

func New() *App {
	catRepo := r.NewCatRepoMemory()
	catHandler := h.NewCatHandler(catRepo)

	shelterRepo := r.NewShelterRepoMemory()
	shelterHandler := h.NewShelterHandler(shelterRepo, catRepo)

	return &App{
		CatHandler:     catHandler,
		ShelterHandler: shelterHandler,
	}
}
