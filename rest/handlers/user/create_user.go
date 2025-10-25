package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/domain"
	"github.com/gowalillah/ecommerce/util"
)

type ResUser struct {
	ID        int    `json:"id"`
	Unique_id string `json:"unique_id"`
	Message   string `json:"message"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	// Role      UserRole  `json:"role" db:"role"`
}

type ReqCreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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

	user, err := h.svc.Create(domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})

	if err != nil {
		http.Error(w, "Internal server err", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, ResUser{
		Message:   "Successfully created user...!",
		ID:        user.ID,
		Unique_id: user.Uuid,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
	})

}
