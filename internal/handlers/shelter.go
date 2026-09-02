package handlers

import "cp_lab1/internal/repositories"

type ShelterHandler struct {
	shelterRepo repositories.ShelterRepo
	catRepo     repositories.CatRepo
}

func NewShelterHandler(sr repositories.ShelterRepo, cr repositories.CatRepo) *ShelterHandler {
	return &ShelterHandler{
		shelterRepo: sr,
		catRepo:     cr,
	}
}
