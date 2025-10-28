package user

import (
	"encoding/json"
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) ChangeUserRole(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID      int    `json:"id"`
		NewRole string `json:"new_role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.UpdateRole(req.ID, req.NewRole); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, map[string]string{"message": "role updated"})
}
