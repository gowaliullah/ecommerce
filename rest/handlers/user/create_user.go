package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

type ReqCreateUser struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	// Role      UserRole  `json:"role" db:"role"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", http.StatusBadRequest)
		return
	}

	createdUser, err := h.svc.Create(domain.User{
		// ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
	})

	if err != nil {
		http.Error(w, "Internal server err", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, createdUser)

}
