package inmemory

import (
	"context"
	"fmt"
	"hw7/internal/models"
	"sync"
)

type BagsRepo struct {
	data map[int]*models.Bag

	mu *sync.RWMutex
}

func (db *BagsRepo) Create(ctx context.Context, Bag *models.Bag) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[Bag.ID] = Bag
	return nil
}

func (db *BagsRepo) All(ctx context.Context) ([]*models.Bag, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	Bags := make([]*models.Bag, 0, len(db.data))
	for _, Bag := range db.data {
		Bags = append(Bags, Bag)
	}

	return Bags, nil
}

func (db *BagsRepo) ByID(ctx context.Context, id int) (*models.Bag, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	Bag, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No Bag with id %d", id)
	}

	return Bag, nil
}

func (db *BagsRepo) Update(ctx context.Context, Bag *models.Bag) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[Bag.ID] = Bag
	return nil
}

func (db *BagsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
