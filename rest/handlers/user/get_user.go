package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	usrId := r.PathValue("id")
	id, err := strconv.Atoi(usrId)

	if err != nil {
		http.Error(w, "Please give the valid user id", http.StatusBadRequest)
		return
	}

	user, err := h.svc.Get(id)
	fmt.Println(user, err)
	if user == nil {
		util.SendError(w, http.StatusNotFound, "User not found")
	}
	util.SendData(w, http.StatusOK, user)
}
