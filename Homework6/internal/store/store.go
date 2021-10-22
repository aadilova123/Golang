package store

import (
	"context"
	"hw6/project/internal/models"
)

type Store interface {
	// Books() BooksRepository
	Create(ctx context.Context, book *models.Book) error
	All(ctx context.Context) ([]*models.Book, error)
	ByID(ctx context.Context, id int) (*models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int) error
}

// type BooksRepository interface {
// 	Create(ctx context.Context, book *models.Book) error
// 	All(ctx context.Context) ([]*models.Book, error)
// 	ByID(ctx context.Context, id string) (*models.Book, error)
// 	Update(ctx context.Context, book *models.Book) error
// 	Delete(ctx context.Context, id string) error
// }
