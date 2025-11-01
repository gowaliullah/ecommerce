package repo

import (
	"github.com/google/uuid"
	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/order"
	"github.com/jmoiron/sqlx"
)

type OrderRepo interface {
	order.OrderRepo
}

type orderRepo struct {
	db sqlx.DB
}

func NewCartItemRepo(db sqlx.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(o domain.Order) (*domain.Order, error) {
	o.Uuid = uuid.New().String()

	query := `
		INSERT INTO orders (
			unique_id,
			user_id,
			product_id,
			status,
			quantity,
			total_price
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		RETURNING id
	`

	row := r.db.QueryRow(
		query,
		o.Uuid,
		o.UserID,
		o.ProductID,
		o.Status,
		o.Quantity,
		o.Total,
	)

	err := row.Scan(&o.ID)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r *orderRepo) List() ([]*domain.Order, error) {
	query := `SELECT id, unique_id, user_id`

	var carts []*domain.Order

	err := r.db.Select(&carts, query)

	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *orderRepo) Get(id int) (*domain.Order, error) {
	query := `
		SELECT id, unique_id, user_id WHERE id = $1
 	`
	var c domain.Order

	err := r.db.Get(&c, query, id)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *orderRepo) Update(c domain.Order) (*domain.Order, error) {
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

func (r *orderRepo) Delete(id int) error {
	query := `DELETE FROM carts WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
