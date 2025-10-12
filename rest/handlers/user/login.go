package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(req.Email, req.Password)

	usr, err := h.svc.Find(req.Email, req.Password)

	fmt.Println("usr: ", usr)
	fmt.Println("err: ", err)

	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr == nil {
		fmt.Println("From here")
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Role:      usr.Role,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(accessToken)
	util.SendData(w, http.StatusOK, accessToken)
}
