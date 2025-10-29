package repo

import (
	"github.com/google/uuid"
	cartitem "github.com/gowalillah/ecommerce/cartitem"
	"github.com/gowalillah/ecommerce/domain"
	"github.com/jmoiron/sqlx"
)

type CartItemRepo interface {
	cartitem.CartItemRepo
}

type cartItemRepo struct {
	db sqlx.DB
}

func NewCartItemRepo(db sqlx.DB) CartItemRepo {
	return &cartItemRepo{
		db: db,
	}
}

func (r *cartItemRepo) Create(c domain.CartItem) (*domain.CartItem, error) {

	c.Uuid = uuid.New().String()

	query := `
		INSERT INTO carts (
			unique_id,
			user_id
		) VALUES (
			$1, $2
		) RETURNING id
	`
	row := r.db.QueryRow(query, c.Uuid, c.CartID, c.ProductID, c.Quantity)
	err := row.Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *cartItemRepo) List() ([]*domain.CartItem, error) {
	query := `SELECT id, unique_id, user_id`

	var carts []*domain.CartItem

	err := r.db.Select(&carts, query)

	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *cartItemRepo) Get(id int) (*domain.CartItem, error) {
	query := `
		SELECT id, unique_id, user_id WHERE id = $1
 	`
	var c domain.CartItem

	err := r.db.Get(&c, query, id)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *cartItemRepo) Update(c domain.CartItem) (*domain.CartItem, error) {
	query := `
		UPDATE carts SET
		user_id = $1,
	WHERE id = $2
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *cartItemRepo) Delete(id int) error {
	query := `DELETE FROM carts WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
