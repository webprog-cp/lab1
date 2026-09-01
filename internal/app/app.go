package app

import h "cp_lab1/internal/handlers"

type App struct {
	CatHandler     *h.CatHandler
	ShelterHandler *h.ShelterHandler
}
