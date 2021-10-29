package inmemory

import (
	"context"
	"fmt"
	"hw7/internal/models"
	"sync"
)

type BraceletsRepo struct {
	data map[int]*models.Bracelet

	mu *sync.RWMutex
}

func (db *BraceletsRepo) Create(ctx context.Context, Bracelet *models.Bracelet) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[Bracelet.ID] = Bracelet
	return nil
}

func (db *BraceletsRepo) All(ctx context.Context) ([]*models.Bracelet, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	Bracelets := make([]*models.Bracelet, 0, len(db.data))
	for _, Bracelet := range db.data {
		Bracelets = append(Bracelets, Bracelet)
	}

	return Bracelets, nil
}

func (db *BraceletsRepo) ByID(ctx context.Context, id int) (*models.Bracelet, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	Bracelet, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No Bracelet with id %d", id)
	}

	return Bracelet, nil
}

func (db *BraceletsRepo) Update(ctx context.Context, Bracelet *models.Bracelet) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[Bracelet.ID] = Bracelet
	return nil
}

func (db *BraceletsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
