package repo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/user"
	"github.com/gowalillah/ecommerce/util"
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

	user.Uuid = uuid.New().String()
	query := `
		INSERT INTO users(unique_id, first_name, last_name, email, password, role) VALUES(
		$1, $2, $3, $4, $5, $6
		)
		RETURNING id
	`

	user.Password = util.CreateHashPassword(user.Password)

	if user.Role == "" {
		user.Role = "user"
	}

	row := r.db.QueryRow(query, user.Uuid, user.FirstName, user.LastName, user.Email, user.Password, user.Role)

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
		SELECT id, unique_id, first_name, last_name, email, password, role
		FROM users
		WHERE email = $1
	`

	err := r.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	// checking the password
	ok, err := util.CheckHashedassword(pass, user.Password)
	if !ok {
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

func (r *userRepo) List(limit, page int) ([]*domain.User, error) {

	if limit <= 0 || limit >= 100 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	var users []*domain.User

	query := `SELECT id, unique_id, first_name, last_name, email, role, created_at, updated_at FROM users LIMIT $1 OFFSET $2`

	err := r.db.Select(&users, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) Get(id string) (*domain.User, error) {

	var usr domain.User

	query := `SELECT id, unique_id, first_name, last_name, email, password, role FROM users WHERE id = $1`
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

func (r *userRepo) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
