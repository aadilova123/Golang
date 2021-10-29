package store

import (
	"context"
	"hw7/internal/models"
)

type Store interface {
	Bags() BagsRepository
	Bracelets() BraceletsRepository
}

type BagsRepository interface {
	Create(ctx context.Context, bag *models.Bag) error
	All(ctx context.Context) ([]*models.Bag, error)
	ByID(ctx context.Context, id int) (*models.Bag, error)
	Update(ctx context.Context, bag *models.Bag) error
	Delete(ctx context.Context, id int) error
}

type BraceletsRepository interface {
	Create(ctx context.Context, bracelet *models.Bracelet) error
	All(ctx context.Context) ([]*models.Bracelet, error)
	ByID(ctx context.Context, id int) (*models.Bracelet, error)
	Update(ctx context.Context, bracelet *models.Bracelet) error
	Delete(ctx context.Context, id int) error
}
