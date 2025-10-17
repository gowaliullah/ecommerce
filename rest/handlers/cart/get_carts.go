package cart

import (
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetCarts(w http.ResponseWriter, r *http.Request) {
	carts, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to retrieve carts")
		return
	}
	util.SendData(w, http.StatusOK, carts)
}
