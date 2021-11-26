package store

import (
	"context"
	"hw8/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Categories() CategoriesRepository
	Accesories() AccesoriesRepository
}

type CategoriesRepository interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
}

type AccesoriesRepository interface {
	Create(ctx context.Context, good *models.Accesory) error
	All(ctx context.Context) ([]*models.Accesory, error)
	ByID(ctx context.Context, id int) (*models.Accesory, error)
	Update(ctx context.Context, good *models.Accesory) error
	Delete(ctx context.Context, id int) error
}