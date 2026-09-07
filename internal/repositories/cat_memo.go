package repositories

import (
	"cp_lab1/internal/const_errors"
	"cp_lab1/internal/models"
	"sync"
)

type CatRepoMemo struct {
	cats   map[uint64]models.Cat
	nextID uint64
	mu     sync.RWMutex
}

func NewCatRepoMemory() *CatRepoMemo {
	return &CatRepoMemo{
		cats:   make(map[uint64]models.Cat),
		nextID: 1,
	}
}

func (cr *CatRepoMemo) GetAll() ([]models.Cat, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()

	cats := make([]models.Cat, 0, len(cr.cats))

	for _, cat := range cr.cats {
		cats = append(cats, cat)
	}

	return cats, nil
}

func (cr *CatRepoMemo) GetByID(id uint64) (models.Cat, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()

	cat, exists := cr.cats[id]
	if !exists {
		return models.Cat{}, const_errors.NoEntityByID
	}

	return cat, nil
}

func (cr *CatRepoMemo) GetCatsByShelterID(id uint64) ([]models.Cat, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()

	var cats []models.Cat

	for _, cat := range cr.cats {
		if cat.ShelterID == id {
			cats = append(cats, cat)
		}
	}

	return cats, nil
}

func (cr *CatRepoMemo) Create(cat models.Cat) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	cat.ID = cr.nextID
	cr.nextID++

	cr.cats[cat.ID] = cat

	return nil
}

func (cr *CatRepoMemo) Update(id uint64, cat models.Cat) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	if _, exists := cr.cats[id]; !exists {
		return const_errors.NoEntityByID
	}

	cat.ID = id
	cr.cats[id] = cat

	return nil
}

func (cr *CatRepoMemo) Delete(id uint64) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	if _, exists := cr.cats[id]; !exists {
		return const_errors.NoEntityByID
	}

	delete(cr.cats, id)

	return nil
}

func (cr *CatRepoMemo) DeleteAll() error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	cr.cats = make(map[uint64]models.Cat)
	cr.nextID = 1

	return nil
}
