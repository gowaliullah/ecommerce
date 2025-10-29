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

	usr, ok := util.GetUserFromContext(r, h.cnf)
	fmt.Println(id, usr, ok)

	// fmt.Println("User not found in context")
	// if !ok {
	// 	util.SendError(w, http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }

	// if usr.ID != id {
	// 	fmt.Println("User trying to access other user's information")
	// 	util.SendError(w, http.StatusForbidden, "You can only access your own information")
	// 	return
	// }

	user, err := h.svc.Get(usrId)
	fmt.Println(user, err)
	if user == nil {
		util.SendError(w, http.StatusNotFound, "User not found")
	}
	util.SendData(w, http.StatusOK, user)
}
