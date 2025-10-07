package product

import (
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", 400)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusOK, "Successfully deleted product")

}
