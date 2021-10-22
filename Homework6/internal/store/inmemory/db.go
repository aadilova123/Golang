package inmemory

import (
	"context"
	"fmt"
	"hw6/project/internal/models"
	"hw6/project/internal/store"	
	"sync"
)

type DB struct {
	data map[int]*models.Book
	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Book),
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Create(ctx context.Context, book *models.Book) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[book.ID] = book
	return nil
}

func (db *DB) All(ctx context.Context) ([]*models.Book, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	books := make([]*models.Book, 0, len(db.data))
	for _, book := range db.data {
		books = append(books, book)
	}

	return books, nil
}

func (db *DB) ByID(ctx context.Context, id int) (*models.Book, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	book, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No product with id %d", id)
	}

	return book, nil
}

func (db *DB) Update(ctx context.Context, book *models.Book) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[book.ID] = book
	return nil
}

func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, id)
	return nil
}