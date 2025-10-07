package product

import (
	"net/http"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}
	util.SendData(w, http.StatusOK, products)
}
