package repo

import (
	"github.com/gowalillah/ecommerce/category"
	"github.com/gowalillah/ecommerce/domain"
	"github.com/jmoiron/sqlx"
)

type CategoryRepo interface {
	category.CategoryRepo
}

type categoryRepo struct {
	db sqlx.DB
}

func NewCategoryRepo(db sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(c domain.Category) (*domain.Category, error) {
	query := `
		INSERT INTO categories (
			name image_url
		) VALUES (
			$1, $2 
		) RETURNING id
	`
	row := r.db.QueryRow(query, c.Name, c.ImageUrl)
	err := row.Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *categoryRepo) List() ([]*domain.Category, error) {
	query := `SELECT id, name, image_url`

	var categories []*domain.Category

	err := r.db.Select(&categories, query)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepo) Get(id int) (*domain.Category, error) {
	query := `
		SELECT id, name, image_url WHERE id = $1
 	`
	var c domain.Category

	err := r.db.Get(&c, query, id)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
func (r *categoryRepo) Update(c domain.Category) (*domain.Category, error) {
	query := `
		UPDATE categories SET 
		name = $1,
		image_url = $2
	WHERE id = $3
	`
	_, err := r.db.Exec(query, c.Name, c.ImageUrl)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
func (r *categoryRepo) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
