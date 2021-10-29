package inmemory

import (
	"hw6/project/internal/models"
	"hw6/project/internal/store"	
	"sync"
)

type DB struct {
	booksRepo store.BooksRepository
	authorsRepo store.AuthorsRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Books() store.BooksRepository {
	if db.booksRepo == nil {
		db.booksRepo = &BooksRepo{
			data: make(map[int]*models.Romance),
			mu: new(sync.RWMutex),
		}
	}
	return db.booksRepo
}

func (db *DB) Authors() store.AuthorsRepository{
	if db.authorsRepo == nil {
		db.authorsRepo = &AuthorsRepo{
			data: make(map[int]*models.Author),
			mu: new(sync.RWMutex),
		}
	}
	return db.authorsRepo
}
