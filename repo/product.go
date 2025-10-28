package repo

import (
	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/product"
	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db sqlx.DB
}

func NewProductRepo(db sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (
			title, description, price, stock, img_url, created_by 
		) VALUES (
		 	$1, $2, $3, $4, $5, $6
		) RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.Stock, p.ImgUrl, p.CreatedBy)

	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM products`

	var count int64

	row := r.db.QueryRow(query).Scan(&count)

	if row != nil {
		return 0, nil
	}

	return count, nil

}

func (r *productRepo) List() ([]*domain.Product, error) {

	query := `SELECT id, title, description, price, stock, img_url, created_by FROM products`

	var products []*domain.Product

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	query := `SELECT id, title, description, price, stock, img_url, created_by FROM products WHERE id = $1`

	var prd domain.Product

	err := r.db.Get(&prd, query, id)
	if err != nil {
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products SET
			title = $1,
			description = $2,
			price = $3,
			stock = $4,
			img_url = $5
		WHERE id = $6
	`

	_, err := r.db.Exec(query, p.Title, p.Description, p.Price, p.Stock, p.ImgUrl, p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
