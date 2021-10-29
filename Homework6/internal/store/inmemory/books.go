package inmemory

import (
	"context"
	"fmt"
	"hw6/project/internal/models"

	"sync"
)

type BooksRepo struct {
	data map[int]*models.Romance
	
	mu sync.RWMutex
}

func (db *BooksRepo) Create(ctx context.Context, book *models.Romance) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[book.ID] = book
	return nil
}

func (db *BooksRepo) All(ctx context.Context) ([]*models.Romance, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	books := make([]*models.Romance, 0, len(db.data))
	for _, book := range db.data {
		books = append(books, book)
	}

	return books, nil
}

func (db *BooksRepo) ByID(ctx context.Context, id int) (*models.Romance, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	book, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No product with id %d", id)
	}

	return book, nil
}

func (db *BooksRepo) Update(ctx context.Context, book *models.Romance) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[book.ID] = book
	return nil
}

func (db *BooksRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, id)
	return nil
}