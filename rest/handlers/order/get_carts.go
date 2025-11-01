package order

import (
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	carts, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.SendData(w, http.StatusOK, carts)
}
