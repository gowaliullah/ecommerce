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

	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr == nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
	// 	Sub:       usr.ID,
	// 	FirstName: usr.FirstName,
	// 	LastName:  usr.LastName,
	// 	Email:     usr.Email,
	// })

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// fmt.Println(accessToken)
	// util.SendData(w, accessToken, http.StatusOK)
}
