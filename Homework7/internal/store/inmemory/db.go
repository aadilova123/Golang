package inmemory

import (
	"hw7/internal/models"
	"hw7/internal/store"
	"sync"
)

type DB struct {
	bagsRepo store.BagsRepository
	braceletsRepo  store.BraceletsRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Bags() store.BagsRepository {
	if db.bagsRepo == nil {
		db.bagsRepo = &BagsRepo{
			data: make(map[int]*models.Bag),
			mu:   new(sync.RWMutex),
		}
	}

	return db.bagsRepo
}

func (db *DB) Bracelets() store.BraceletsRepository {
	if db.braceletsRepo == nil {
		db.braceletsRepo = &BraceletsRepo{
			data: make(map[int]*models.Bracelet),
			mu:   new(sync.RWMutex),
		}
	}

	return db.braceletsRepo
}