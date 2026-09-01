package repositories

import "cp_lab1/internal/models"

type ShelterRepo interface {
	GetAll() []models.Cat
	GetByID() (models.Cat, error)
}
