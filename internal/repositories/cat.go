package repositories

import "cp_lab1/internal/models"

type CatRepo interface {
	GetAll() []models.Cat
	GetByID(uint64) (models.Cat, error)
	Create(models.Cat) (models.Cat, error)
	Update(uint64, models.Cat) error
	Delete(uint64) error
	DeleteAll()
}
