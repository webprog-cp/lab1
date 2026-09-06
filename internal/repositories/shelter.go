package repositories

import "cp_lab1/internal/models"

type ShelterRepo interface {
	GetAll() []models.Shelter
	GetByID(uint64) (models.Shelter, error)
	Create(models.Shelter) (models.Shelter, error)
	Update(uint64, models.Shelter) error
	Delete(uint64) error
	DeleteAll() error
}
