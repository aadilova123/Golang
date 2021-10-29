package inmemory

import (
	"context"
	"fmt"
	"hw6/project/internal/models"
	"sync"
)

type AuthorsRepo struct {
	data map[int]*models.Author
	mu sync.RWMutex
}

func (db *AuthorsRepo) Create(ctx context.Context, a *models.Author) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[a.ID] = a
	return nil
}

func (db *AuthorsRepo) All(ctx context.Context) ([]*models.Author, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	authors := make([]*models.Author, 0, len(db.data))
	for _, author := range db.data {
		authors = append(authors, author)
	}

	return authors, nil
}

func (db *AuthorsRepo) ByID(ctx context.Context, id int) (*models.Author, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	author, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No product with id %d", id)
	}

	return author, nil
}

func (db *AuthorsRepo) Update(ctx context.Context, author *models.Author) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[author.ID] = author
	return nil
}

func (db *AuthorsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, id)
	return nil
}