package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

type ReqLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResLogin struct {
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid credentials", http.StatusBadRequest)
		return
	}

	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusNotFound, "Email or Password is incorrect")
		return
	}

	if usr == nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	fmt.Println(usr.Uuid)

	accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
		ID:        usr.ID,
		Uuid:      usr.Uuid,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Role:      usr.Role,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, ResLogin{
		Message:     "Successfully user login..!",
		AccessToken: accessToken,
	})
}
