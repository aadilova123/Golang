package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"hw8/internal/models"
	"hw8/internal/store"
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

func (c GoodsRepository) Create(ctx context.Context, category *models.Good) error {
	_, err := c.conn.Exec("INSERT INTO goods(name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c GoodsRepository) All(ctx context.Context) ([]*models.Good, error) {
	goods := make([]*models.Good, 0)
	if err := c.conn.Select(&goods, "SELECT * FROM goods"); err != nil {
		return nil, err
	}

	return goods, nil
}

func (c GoodsRepository) ByID(ctx context.Context, id int) (*models.Good, error) {
	good := new(models.Good)
	if err := c.conn.Get(good, "SELECT id, name FROM goods WHERE id=$1", id); err != nil {
		return nil, err
	}

	return good, nil
}

func (c GoodsRepository) Update(ctx context.Context, category *models.Good) error {
	_, err := c.conn.Exec("UPDATE goods SET name = $1 WHERE id = $2", category.Name, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c GoodsRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM goods WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}