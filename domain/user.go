package domain

import "time"

// UserRole type (string) for ENUM mapping
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	Uuid      string `json:"unique_id" db:"unique_id" validate:"required,unique"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email" validate:"required,email"`
	Password  string `json:"password" db:"password" validate:"password"`
	// Role      UserRole  `json:"role" db:"role"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
