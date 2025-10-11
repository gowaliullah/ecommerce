package user

import (
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}
	util.SendData(w, http.StatusOK, users)
}
