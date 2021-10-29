package store

import (
	"context"
	"hw6/project/internal/models"
)

type Store interface {
	Books() BooksRepository
	Authors() AuthorsRepository
}

type BooksRepository interface {
	Create(ctx context.Context, book *models.Romance) error
	All(ctx context.Context) ([]*models.Romance, error)
	ByID(ctx context.Context, id int) (*models.Romance, error)
	Update(ctx context.Context, book *models.Romance) error
	Delete(ctx context.Context, id int) error
}

type AuthorsRepository interface {
	Create(ctx context.Context, author *models.Author) error
	All(ctx context.Context) ([]*models.Author, error)
	ByID(ctx context.Context, id int) (*models.Author, error)
	Update(ctx context.Context, author *models.Author) error
	Delete(ctx context.Context, id int) error
}
