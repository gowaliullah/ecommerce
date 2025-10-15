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
func (r *categoryRepo)  {
	
}
func (r *categoryRepo)  {
	
}
func (r *categoryRepo)  {
	
}
func (r *categoryRepo)  {
	
}
func (r *categoryRepo)  {
	
}