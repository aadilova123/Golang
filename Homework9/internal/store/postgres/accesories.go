package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"hw8/internal/models"
	"hw8/internal/store"
)

func (db *DB) Accesories() store.AccesoriesRepository {
	if db.goods == nil {
		db.goods = NewAccesoriesRepository(db.conn)
	}

	return db.goods
}

type AccesoriesRepository struct {
	conn *sqlx.DB
}

func NewAccesoriesRepository(conn *sqlx.DB) store.AccesoriesRepository {
	return &AccesoriesRepository{conn: conn}
}

func (c AccesoriesRepository) Create(ctx context.Context, reciept *models.Accesory) error {
	_, err := c.conn.Exec("INSERT INTO goods(name, category_id) VALUES ($1, $2)", reciept.Name, reciept.CategoryID)
	if err != nil {
		return err
	}

	return nil
}

func (c AccesoriesRepository) All(ctx context.Context) ([]*models.Accesory, error) {
	reciepts := make([]*models.Accesory, 0)
	if err := c.conn.Select(&reciepts, "SELECT * FROM goods"); err != nil {
		return nil, err
	}

	return reciepts, nil
}

func (c AccesoriesRepository) ByID(ctx context.Context, id int) (*models.Accesory, error) {
	reciept := new(models.Accesory)
	if err := c.conn.Get(reciept, "SELECT * FROM goods WHERE id=$1", id); err != nil {
		return nil, err
	}

	return reciept, nil
}

func (c AccesoriesRepository) Update(ctx context.Context, reciept *models.Accesory) error {
	_, err := c.conn.Exec("UPDATE goods SET name = $1  WHERE id = $4", reciept.Name, reciept.ID)
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