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
			title, description, price, stock, img_url
		) VALUES (
		 	$1, $2, $3, $4, $5
		) RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.Stock, p.ImgUrl)

	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
