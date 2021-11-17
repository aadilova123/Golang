package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"hw7/internal/models"
	"hw7/internal/store"
)
	

func (db *DB) Goods() store.GoodsRepository {
	if db.goods == nil {
		db.goods = NewGoodsRepository(db.conn)
	}

	return db.goods
}

type GoodsRepository struct {
	conn *sqlx.DB
}

func NewGoodsRepository(conn *sqlx.DB) store.GoodsRepository {
	return &GoodsRepository{conn: conn}
}

func (c GoodsRepository) Create(ctx context.Context, category *models.Category) error {
	_, err := c.conn.Exec("INSERT INTO Goods(name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c GoodsRepository) All(ctx context.Context) ([]*models.Category, error) {
	Goods := make([]*models.Category, 0)
	if err := c.conn.Select(&Goods, "SELECT * FROM Goods"); err != nil {
		return nil, err
	}

	return Goods, nil
}

func (c GoodsRepository) ByID(ctx context.Context, id int) (*models.Category, error) {
	category := new(models.Category)
	if err := c.conn.Get(category, "SELECT id, name FROM Goods WHERE id=$1", id); err != nil {
		return nil, err
	}

	return category, nil
}

func (c GoodsRepository) Update(ctx context.Context, category *models.Category) error {
	_, err := c.conn.Exec("UPDATE Goods SET name = $1 WHERE id = $2", category.Name, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c GoodsRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM Goods WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}