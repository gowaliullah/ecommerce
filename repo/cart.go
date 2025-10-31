package repo

import (
	"github.com/google/uuid"
	"github.com/gowalillah/ecommerce/cart"
	"github.com/gowalillah/ecommerce/domain"
	"github.com/jmoiron/sqlx"
)

type CartRepo interface {
	cart.CartRepo
}

type cartRepo struct {
	db sqlx.DB
}

func NewCartRepo(db sqlx.DB) CartRepo {
	return &cartRepo{
		db: db,
	}
}

func (r *cartRepo) Create(c domain.Cart) (*domain.Cart, error) {
	c.Uuid = uuid.New().String()

	query := `
		INSERT INTO carts (
			unique_id,
			user_id
		) VALUES ($1, $2)
		RETURNING id
	`

	var userID interface{}
	if c.UserID == "" {
		userID = nil
	} else {
		userID = c.UserID
	}

	row := r.db.QueryRow(query, c.Uuid, userID)
	err := row.Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *cartRepo) List() ([]*domain.Cart, error) {
	query := `SELECT id, unique_id, user_id FROM carts`

	var carts []*domain.Cart

	err := r.db.Select(&carts, query)

	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *cartRepo) Get(id int) (*domain.Cart, error) {
	query := `
		SELECT id, unique_id, user_id FROM carts WHERE id = $1
 	`
	var c domain.Cart

	err := r.db.Get(&c, query, id)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *cartRepo) Update(c domain.Cart) (*domain.Cart, error) {
	query := `
		UPDATE carts SET
		user_id = $1,
	WHERE id = $2
	`
	_, err := r.db.Exec(query, c.UserID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *cartRepo) Delete(id int) error {
	query := `DELETE FROM carts WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
