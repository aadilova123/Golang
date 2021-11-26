package postgres

import (
	"context"
	"hw8/internal/models"
	"hw8/internal/store"

	"github.com/jmoiron/sqlx"
)

func (db *DB) Accesories() store.AccesoriesRepository {
	if db.Accesories() == nil {
		db.reciepts = NewAccesoriesRepository(db.conn)
	}
	return db.reciepts
}

type AccesoriesRepository struct {
	conn *sqlx.DB
}

func NewAccesoriesRepository(conn *sqlx.DB) store.AccesoriesRepository {
	return &AccesoriesRepository{conn: conn}
}

func (c AccesoriesRepository) Create(ctx context.Context, reciept *models.Accesory) error {
	_, err := c.conn.Exec("INSERT INTO accesories(name, category_id, brand, material) VALUES ($1, $2, $3, $4)", reciept.Name, reciept.CategoryID, reciept.Brand, reciept.Material)
	if err != nil {
		return err
	}

	return nil
}

func (c AccesoriesRepository) All(ctx context.Context) ([]*models.Accesory, error) {
	reciepts := make([]*models.Accesory, 0)
	if err := c.conn.Select(&reciepts, "SELECT * FROM reciepts"); err != nil {
		return nil, err
	}

	return reciepts, nil
}

func (c AccesoriesRepository) ByID(ctx context.Context, id int) (*models.Accesory, error) {
	reciept := new(models.Accesory)
	if err := c.conn.Get(reciept, "SELECT * FROM categories WHERE id=$1", id); err != nil {
		return nil, err
	}

	return reciept, nil
}

func (c AccesoriesRepository) Update(ctx context.Context, reciept *models.Accesory) error {
	_, err := c.conn.Exec("UPDATE jobs SET name = $1, brand = $2, material = $3  WHERE id = $4", reciept.Name, reciept.Brand, reciept.Material, reciept.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c AccesoriesRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM accesories WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}