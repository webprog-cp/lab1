package repositories

import (
	"cp_lab1/internal/errors"
	"cp_lab1/internal/models"
	"sync"
)

type ShelterRepoMemo struct {
	shelters map[uint64]models.Shelter
	nextID   uint64
	mu       sync.RWMutex
}

func NewShelterRepoMemory() *ShelterRepoMemo {
	return &ShelterRepoMemo{
		shelters: make(map[uint64]models.Shelter),
		nextID:   1,
	}
}

func (sr *ShelterRepoMemo) GetAll() []models.Shelter {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	shelters := make([]models.Shelter, 0, len(sr.shelters))

	for _, shelter := range sr.shelters {
		shelters = append(shelters, shelter)
	}

	return shelters
}

func (sr *ShelterRepoMemo) GetByID(id uint64) (models.Shelter, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	shelter, exists := sr.shelters[id]
	if !exists {
		return models.Shelter{}, errors.NoEntityByID
	}

	return shelter, nil
}

func (sr *ShelterRepoMemo) Create(shelter models.Shelter) (models.Shelter, error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	shelter.ID = sr.nextID
	sr.nextID++

	sr.shelters[shelter.ID] = shelter

	return shelter, nil
}

func (sr *ShelterRepoMemo) Update(id uint64, shelter models.Shelter) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	if _, exists := sr.shelters[id]; !exists {
		return errors.NoEntityByID
	}

	shelter.ID = id
	sr.shelters[id] = shelter

	return nil
}

func (sr *ShelterRepoMemo) Delete(id uint64) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	if _, exists := sr.shelters[id]; !exists {
		return errors.NoEntityByID
	}

	delete(sr.shelters, id)

	return nil
}

func (sr *ShelterRepoMemo) DeleteAll() {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.shelters = make(map[uint64]models.Shelter)
	sr.nextID = 1
}
