package cart

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.PathValue("id")
	id, err := strconv.Atoi(cartId)
	if err != nil {
		http.Error(w, "Please give the valid cart id", http.StatusBadRequest)
		return
	}

	cart, _ := h.svc.Get(id)
	if cart == nil {
		util.SendError(w, http.StatusNotFound, "cart not found")
	}

	util.SendData(w, http.StatusOK, cart)
}
