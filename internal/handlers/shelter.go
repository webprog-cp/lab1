package handlers

import "cp_lab1/internal/repositories"

type ShelterHandler struct {
	repo repositories.ShelterRepo
}

func NewShelterHandler(repo repositories.ShelterRepo) *ShelterHandler {
	return &ShelterHandler{
		repo: repo,
	}
}
