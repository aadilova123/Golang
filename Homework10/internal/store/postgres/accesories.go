package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hw10/internal/models"
	"hw10/internal/store"
)

func (db *DB) Accesories() store.AccesoriesRepository {
	if db.goods == nil {
		db.goods = NewAccesoryRepository(db.conn)
	}

	return db.goods
}



type AccesoriesRepository struct {
	conn *sqlx.DB
}

func NewAccesoryRepository(conn *sqlx.DB) store.AccesoriesRepository {
	return &AccesoriesRepository{conn: conn}
}

func (c AccesoriesRepository) Create(ctx context.Context, category *models.Accesory) error {
	_, err := c.conn.Exec("INSERT INTO goods(name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c AccesoriesRepository) All(ctx context.Context, filter *models.AccesoryFilter) ([]*models.Accesory, error) {
	categories := make([]*models.Accesory, 0)
	basicQuery := "SELECT * FROM goods"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE name ILIKE $1", basicQuery)

		if err := c.conn.Select(&categories, basicQuery, "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}

		return categories, nil
	}

	if err := c.conn.Select(&categories, basicQuery); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c AccesoriesRepository) ByID(ctx context.Context, id int) (*models.Accesory, error) {
	category := new(models.Accesory)
	if err := c.conn.Get(category, "SELECT id, name FROM goods WHERE id=$1", id); err != nil {
		return nil, err
	}

	return category, nil
}

func (c AccesoriesRepository) ByCategoryID(ctx context.Context, id int) (*models.Accesory, error) {
	category := new(models.Accesory)
	if err := c.conn.Get(category, "SELECT category_id, name FROM goods WHERE category_id=$1", id); err != nil {
		return nil, err
	}

	return category, nil
}

func (c AccesoriesRepository) Update(ctx context.Context, category *models.Accesory) error {
	_, err := c.conn.Exec("UPDATE goods SET name = $1 WHERE id = $2", category.Name, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c AccesoriesRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM goods WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}