package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/user"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {

	query := `
		INSERT INTO users(first_name, last_name, email, password, role) VALUES(
		$1, $2, $3, $4, $5
		)
		RETURNING id
	`

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role)

	if row.Err() != nil {
		fmt.Println("err", row.Err())
		return nil, row.Err()
	}

	row.Scan(&user.ID)

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, first_name, last_name, email, password, role
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`

	fmt.Println("repo: ", email, pass)
	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int64

	row := r.db.QueryRow(query).Scan(&count)

	if row != nil {
		return 0, nil
	}

	return count, nil

}

func (r *userRepo) List() ([]*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, password, role FROM users`

	var users []*domain.User

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) Get(id int) (*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, password, role FROM users WHERE id = $1`

	var usr domain.User

	err := r.db.Get(&usr, query, id)
	if err != nil {
		return nil, err
	}

	return &usr, nil

}

func (r *userRepo) Update(u domain.User) (*domain.User, error) {
	query := `
		UPDATE users SET 
			first_name = $1, 
			last_name = $2, 
			email = $3, 
			password = $4, 
			role = $5
		WHERE id = $6
	`
	_, err := r.db.Exec(query, u.FirstName, u.LastName, u.Email, u.Password, u.Role, u.Role)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
